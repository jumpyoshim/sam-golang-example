package repo

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
)

func BatchWrite(
	ctx context.Context,
	svc dynamodbiface.DynamoDBAPI,
	in []user.User,
) (out *dynamodb.BatchWriteItemOutput, err error) {
	var (
		reqs []*dynamodb.WriteRequest
	)

	for _, v := range in {
		item, err := dynamodbattribute.MarshalMap(v)
		if err != nil {
			return out, err
		}
		reqs = append(reqs, &dynamodb.WriteRequest{PutRequest: &dynamodb.PutRequest{Item: item}})
	}

	out, err = svc.BatchWriteItemWithContext(ctx, &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{user.TableName: reqs},
	})
	if err != nil {
		return out, err
	}

	return out, err
}
