package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"app/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// These run against a real DynamoDB, because what is being tested is DynamoDB's atomicity and a
// fake would only be testing itself. Point DYNAMO_ENDPOINT at a local instance:
//
//	podman run -d -p 8000:8000 docker.io/amazon/dynamodb-local
//	DYNAMO_ENDPOINT=http://localhost:8000 go test ./db/...
func requireDynamo(t *testing.T) {
	t.Helper()
	if len(os.Getenv("DYNAMO_ENDPOINT")) == 0 {
		t.Skip("DYNAMO_ENDPOINT no está definido; se omite la prueba de integración")
	}

	core.LoadEnv()
	core.Env.DYNAMO_TABLE = "unicore-test-db"

	client, err := Client()
	if err != nil {
		t.Fatalf("no se pudo crear el cliente: %v", err)
	}

	_, err = client.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName:   aws.String(core.Env.DYNAMO_TABLE),
		BillingMode: types.BillingModePayPerRequest,
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("pk"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("sk"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("pk"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("sk"), KeyType: types.KeyTypeRange},
		},
	})
	var alreadyExists *types.ResourceInUseException
	if err != nil && !errors.As(err, &alreadyExists) {
		t.Fatalf("no se pudo crear la tabla: %v", err)
	}
}

// Each test needs an identifier nobody else is counting against, or a rerun inside the same
// window would start with the budget already spent.
var identifierCounter atomic.Int64

func freshIdentifier() int64 {
	return time.Now().UnixNano() + identifierCounter.Add(1)
}

func TestCheckAndCountSpendsExactlyTheBudget(t *testing.T) {
	requireDynamo(t)
	ctx := context.Background()
	identifier := freshIdentifier()

	for attempt := int64(1); attempt <= 3; attempt++ {
		decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, 3, time.Minute)
		if err != nil {
			t.Fatalf("intento %v devolvió error: %v", attempt, err)
		}
		if !decision.Allowed {
			t.Fatalf("el intento %v de 3 fue rechazado", attempt)
		}
		if decision.Hits != attempt {
			t.Fatalf("el intento %v reportó %v consumos", attempt, decision.Hits)
		}
	}

	decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, 3, time.Minute)
	if err != nil {
		t.Fatalf("el rechazo no debe ser un error: %v", err)
	}
	if decision.Allowed {
		t.Fatal("el cuarto intento debería haber sido rechazado")
	}
	if decision.RetryAfter <= 0 || decision.RetryAfter > time.Minute {
		t.Fatalf("Retry-After fuera de la ventana: %v", decision.RetryAfter)
	}

	// A refused attempt must not count. Otherwise a caller hammering a blocked key extends its
	// own lockout, and with it the lockout of everyone sharing that NAT.
	if decision.Hits != 3 {
		t.Fatalf("el rechazo movió el contador a %v", decision.Hits)
	}
}

// The reason this design exists. Concurrent callers must not all read "under the limit" and all
// proceed — the exact race the genix backend needs a distributed lock to prevent, and that a
// conditional UpdateItem prevents on its own.
func TestCheckAndCountIsAtomicUnderConcurrency(t *testing.T) {
	requireDynamo(t)
	ctx := context.Background()
	identifier := freshIdentifier()

	const limit = int64(5)
	const callers = 40

	allowed := atomic.Int64{}
	failures := make(chan error, callers)
	seenHits := sync.Map{}

	waitGroup := sync.WaitGroup{}
	start := make(chan struct{})
	for caller := 0; caller < callers; caller++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			<-start
			decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, limit, time.Minute)
			if err != nil {
				failures <- err
				return
			}
			if decision.Allowed {
				allowed.Add(1)
				// Every winner must be handed a different counter value: that number becomes the
				// tail of the message's sort key, so a duplicate would be one message silently
				// overwriting another.
				if _, duplicated := seenHits.LoadOrStore(decision.Hits, true); duplicated {
					failures <- fmt.Errorf("dos llamadas recibieron el mismo contador: %v", decision.Hits)
				}
			}
		}()
	}
	close(start)
	waitGroup.Wait()
	close(failures)

	for err := range failures {
		t.Fatalf("fallo en una llamada concurrente: %v", err)
	}
	if allowed.Load() != limit {
		t.Fatalf("%v de %v llamadas concurrentes pasaron, se esperaban %v",
			allowed.Load(), callers, limit)
	}
}

// Two actions keyed on the same number are two different budgets, which is what lets a future
// feature reuse an IP key without eating the contact form's allowance.
func TestActionsDoNotShareABudget(t *testing.T) {
	requireDynamo(t)
	ctx := context.Background()
	identifier := freshIdentifier()

	const otherAction = core.LimitAction(9999)

	if decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, 1, time.Minute); err != nil || !decision.Allowed {
		t.Fatalf("el primer consumo debería pasar: %v %v", decision.Allowed, err)
	}
	if decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, 1, time.Minute); err != nil || decision.Allowed {
		t.Fatalf("el segundo consumo de la misma acción debería fallar: %v %v", decision.Allowed, err)
	}
	if decision, err := CheckAndCount(ctx, otherAction, identifier, 1, time.Minute); err != nil || !decision.Allowed {
		t.Fatalf("otra acción no debería estar agotada: %v %v", decision.Allowed, err)
	}
}

// A window that has rolled over hands back a full budget. Driven with a one-second window so the
// test does not have to wait out a real one.
func TestBudgetResetsWithTheWindow(t *testing.T) {
	requireDynamo(t)
	ctx := context.Background()
	identifier := freshIdentifier()

	if decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, 1, time.Second); err != nil || !decision.Allowed {
		t.Fatalf("el primer consumo debería pasar: %v %v", decision.Allowed, err)
	}

	// Sleeping to the far side of the next boundary, rather than one second, because the window
	// is aligned to wall-clock time and the first call may have landed anywhere inside it.
	time.Sleep(1100 * time.Millisecond)
	deadline := time.Now().Add(3 * time.Second)
	for {
		decision, err := CheckAndCount(ctx, core.ActionContactByIP, identifier, 1, time.Second)
		if err != nil {
			t.Fatalf("error tras el cambio de ventana: %v", err)
		}
		if decision.Allowed {
			return
		}
		if time.Now().After(deadline) {
			t.Fatal("la ventana nunca se reinició")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
