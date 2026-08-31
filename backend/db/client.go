package db

import (
	"context"
	"fmt"
	"sync"

	"app/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// One lazily built client, shared for the life of the process. The AWS SDK client is
// goroutine-safe and holds the connection pool, so building it per request would pay a TLS
// handshake on every invocation — in Lambda the container outlives the request, which is exactly
// what makes a package-level client worth having.
var (
	clientOnce sync.Once
	clientRef  *dynamodb.Client
	clientErr  error
)

func Client() (*dynamodb.Client, error) {
	clientOnce.Do(func() {
		awsCfg, err := awsconfig.LoadDefaultConfig(context.Background())
		if err != nil {
			clientErr = fmt.Errorf("error al cargar la configuración de AWS: %w", err)
			return
		}
		options := []func(*dynamodb.Options){}
		if len(core.Env.DYNAMO_ENDPOINT) > 0 {
			options = append(options, func(o *dynamodb.Options) {
				o.BaseEndpoint = aws.String(core.Env.DYNAMO_ENDPOINT)
			})
		}
		clientRef = dynamodb.NewFromConfig(awsCfg, options...)
	})
	return clientRef, clientErr
}

// A single table holds every entity, distinguished by a prefix on the partition key. It is one
// table and not three because DynamoDB charges and scales per table, and because nothing here
// ever needs to read two entities together — the prefix is the whole schema.
func tableName() string { return core.Env.DYNAMO_TABLE }

func stringValue(value string) types.AttributeValue {
	return &types.AttributeValueMemberS{Value: value}
}

func numberValue[T int | int32 | int64](value T) types.AttributeValue {
	return &types.AttributeValueMemberN{Value: fmt.Sprint(value)}
}
