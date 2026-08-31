package db

import (
	"context"
	"fmt"
	"time"

	"app/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	// ContactStatusDelivered means the row was stored and the notification reached the inbox.
	ContactStatusDelivered = int32(1)
	// ContactStatusUndelivered means the row was stored but the notification failed. The message
	// is not lost, it just has not been announced — which is a thing an operator can act on, and
	// the reason the two are different numbers rather than a boolean nobody would set.
	ContactStatusUndelivered = int32(2)
)

// ContactMessage is one submission of the public contact form.
type ContactMessage struct {
	// WeekCode is core.WeekCodeAt of the moment it arrived: the partition.
	WeekCode int32
	// IP is core.ClientIPKey of the sender. It leads the sort key rather than sitting in an index
	// because the one question anybody will ask of this table beyond "read the newest" is "what
	// else did this address send", and as a sort-key prefix that is a range query over exactly
	// those items instead of a scan of the week.
	IP int64
	// Created is a Unix timestamp in seconds.
	Created int64
	// Sequence is the rate limiter's counter for this sender inside the current window. It is the
	// tail of the sort key, and it is what keeps two messages the same IP sends inside one second
	// as two distinct rows: the limiter's ADD is atomic, so no two concurrent writers can ever be
	// handed the same value for the same key.
	Sequence int64

	Name    string
	Email   string
	Company string
	Message string
	Status  int32
	Updated int64
}

func contactPartitionKey(weekCode int32) string { return fmt.Sprintf("MSG#%d", weekCode) }

// Zero-padded on every component so the sort order is (address, then time, then arrival) rather
// than the lexicographic accident of unpadded decimals, where "10" sorts before "9".
func contactSortKey(ip int64, created int64, sequence int64) string {
	return fmt.Sprintf("%019d#%011d#%06d", ip, created, sequence)
}

// InsertContactMessage stores one message. It refuses to overwrite an existing item rather than
// silently replacing it, so a key collision surfaces as an error instead of as a message that
// quietly disappeared.
func InsertContactMessage(ctx context.Context, message *ContactMessage) error {
	client, err := Client()
	if err != nil {
		return err
	}

	item := map[string]types.AttributeValue{
		"pk":      stringValue(contactPartitionKey(message.WeekCode)),
		"sk":      stringValue(contactSortKey(message.IP, message.Created, message.Sequence)),
		"ip":      numberValue(message.IP),
		"created": numberValue(message.Created),
		"name":    stringValue(message.Name),
		"email":   stringValue(message.Email),
		"message": stringValue(message.Message),
		"status":  numberValue(message.Status),
		"upd":     numberValue(message.Updated),
	}
	// An empty string is a legal DynamoDB value, but writing one for an optional field that was
	// left blank stores a fact nobody asserted. The attribute is simply absent instead.
	if len(message.Company) > 0 {
		item["company"] = stringValue(message.Company)
	}
	if core.Env.CONTACT_TTL_DAYS > 0 {
		expiresAt := message.Created + core.Env.CONTACT_TTL_DAYS*24*60*60
		item["ttl"] = numberValue(expiresAt)
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(tableName()),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(pk)"),
	})
	if err != nil {
		return fmt.Errorf("error al registrar el mensaje de contacto: %w", err)
	}
	return nil
}

// MarkContactMessageUndelivered records that the notification never reached the inbox. It is a
// separate write and not part of the insert because the order matters: the row is written before
// the mail is attempted, so a broken mail server costs the notification and never the message.
func MarkContactMessageUndelivered(ctx context.Context, message *ContactMessage) error {
	client, err := Client()
	if err != nil {
		return err
	}

	_, err = client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName()),
		Key: map[string]types.AttributeValue{
			"pk": stringValue(contactPartitionKey(message.WeekCode)),
			"sk": stringValue(contactSortKey(message.IP, message.Created, message.Sequence)),
		},
		UpdateExpression: aws.String("SET #status = :status, upd = :upd"),
		// "status" is a DynamoDB reserved word and cannot appear literally in an expression.
		ExpressionAttributeNames: map[string]string{"#status": "status"},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":status": numberValue(ContactStatusUndelivered),
			":upd":    numberValue(time.Now().Unix()),
		},
	})
	if err != nil {
		return fmt.Errorf("error al marcar el mensaje como no entregado: %w", err)
	}
	return nil
}
