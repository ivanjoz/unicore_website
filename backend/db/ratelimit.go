package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"app/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Per-identifier rate limiting, in one conditional write.
//
// The problem it solves is the same one a distributed lock solves: concurrent Lambdas serving the
// same client read the same counter, all conclude they are under the limit, and all write. Read
// the count and then decide, and the limit is worth nothing under exactly the traffic it exists
// to stop — parallel requests from one source are the abuse pattern, not the exception.
//
// DynamoDB answers that without a lock. UpdateItem applies its ConditionExpression and its ADD as
// one atomic operation on one item, so "increment only if still under the limit" is a single
// indivisible step: of N simultaneous callers, exactly the first `limit` of them succeed and the
// rest come back with ConditionalCheckFailed. That is one round trip, no lock to acquire, no lock
// to release, and nothing left holding a key if the Lambda is killed mid-request.
//
// The window is fixed, not sliding: every identifier's counter resets on the same wall-clock
// boundary. The cost is a burst — a caller can spend its whole budget at the end of one window
// and again at the start of the next, so 2×limit inside one window's length is reachable across a
// boundary. That is accepted here because the alternative that keeps a true sliding window is to
// store one item per request and sum them, which is a read followed by a write, which is the race
// this design exists to avoid. For a contact form, a doubled burst at a boundary is not the
// failure mode worth reintroducing a lock for; set the window from how much mail one visitor may
// send in a stretch, and read the limit as "at most 2×limit per window in the worst case".
type LimitDecision struct {
	// Allowed is whether the caller may proceed. It is the only field a handler must read.
	Allowed bool
	// Hits is the counter after this request, so the first allowed call reports 1. Left at limit
	// when the request is refused, because the refused attempt is deliberately not counted: an
	// attacker hammering a blocked key must not be able to extend their own lockout, which would
	// turn the limiter into a way of punishing the legitimate visitor behind the same NAT.
	Hits int64
	// Limit is what the decision was made against, so a caller can build its message without
	// re-reading the environment.
	Limit int64
	// RetryAfter is how long until the window rolls over. Sent as the Retry-After header, which
	// is what makes a 429 actionable instead of just a wall.
	RetryAfter time.Duration
}

// CheckAndCount consumes one unit of the budget for (action, identifier) and reports whether the
// caller may proceed.
//
// A returned error is an infrastructure failure, never a refusal — refusals arrive as
// Allowed=false with a nil error. Callers must fail closed on the error: the endpoints behind
// this limiter are unauthenticated and cost money to serve, so a DynamoDB outage that made the
// limit unenforceable would otherwise turn them into open relays.
//
// The counter item is written with a TTL, so nothing has to be cleaned up: DynamoDB deletes
// expired windows on its own, for free, and a window that lingers past its expiry is harmless
// because the key already names the window it belongs to.
func CheckAndCount(
	ctx context.Context, action core.LimitAction, identifier int64, limit int64, window time.Duration,
) (LimitDecision, error) {

	decision := LimitDecision{Limit: limit}
	if limit <= 0 || window <= 0 {
		return decision, fmt.Errorf("límite (%v) y ventana (%v) deben ser positivos", limit, window)
	}

	client, err := Client()
	if err != nil {
		return decision, err
	}

	windowSeconds := int64(window / time.Second)
	if windowSeconds < 1 {
		windowSeconds = 1
	}
	now := time.Now().Unix()
	windowStart := now - (now % windowSeconds)
	windowEnd := windowStart + windowSeconds
	decision.RetryAfter = time.Duration(windowEnd-now) * time.Second

	// Two windows of grace before the item is eligible for deletion. DynamoDB's TTL sweeper is
	// asynchronous and documented to run within days, so the expiry time is a floor and never a
	// guarantee; the grace only makes sure the sweeper can never remove a window that is still
	// being counted against, which would hand a caller a fresh budget mid-window.
	expiresAt := windowEnd + windowSeconds*2

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName()),
		Key: map[string]types.AttributeValue{
			"pk": stringValue(rateLimitPartitionKey(action, identifier)),
			"sk": stringValue(rateLimitSortKey(windowStart)),
		},
		UpdateExpression:    aws.String("SET #ttl = :ttl ADD #hits :one"),
		ConditionExpression: aws.String("attribute_not_exists(#hits) OR #hits < :limit"),
		ExpressionAttributeNames: map[string]string{
			"#hits": "hits",
			"#ttl":  "ttl",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":one":   numberValue(int64(1)),
			":limit": numberValue(limit),
			":ttl":   numberValue(expiresAt),
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	}

	output, err := client.UpdateItem(ctx, input)
	if err != nil {
		var conditionFailed *types.ConditionalCheckFailedException
		if errors.As(err, &conditionFailed) {
			// The budget is spent. Not an error: this is the limiter working.
			decision.Hits = limit
			return decision, nil
		}
		return decision, fmt.Errorf("error al aplicar el límite de tasa: %w", err)
	}

	decision.Allowed = true
	decision.Hits = readNumber(output.Attributes["hits"], 1)
	return decision, nil
}

// The action leads the key so two features never share a counter even when they key on the same
// number, and so one feature's traffic never lands in another's partition.
func rateLimitPartitionKey(action core.LimitAction, identifier int64) string {
	return fmt.Sprintf("RL#%d#%d", action, identifier)
}

// Zero-padded so the windows of one key sort in time order. Nothing reads them in a range today,
// but an unordered sort key is the kind of thing that cannot be changed later without a migration.
func rateLimitSortKey(windowStart int64) string {
	return fmt.Sprintf("W#%011d", windowStart)
}

func readNumber(value types.AttributeValue, fallback int64) int64 {
	asNumber, ok := value.(*types.AttributeValueMemberN)
	if !ok {
		return fallback
	}
	parsed, err := strconv.ParseInt(asNumber.Value, 10, 64)
	if err != nil {
		return fallback
	}
	return parsed
}
