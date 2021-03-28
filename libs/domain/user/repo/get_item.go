package repo

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
)

func GetItem(
	ctx context.Context,
	svc dynamodbiface.DynamoDBAPI,
	in user.UserKey,
) (out *user.User, err error) {
	key, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return out, err
	}

	o, err := svc.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(user.TableName),
		Key:       key,
	})
	if err != nil {
		return out, err
	}

	if o == nil || len(o.Item) == 0 {
		return out, nil
	}

	err = dynamodbattribute.UnmarshalMap(o.Item, &out)
	if err != nil {
		return out, err
	}

	return out, err
}
