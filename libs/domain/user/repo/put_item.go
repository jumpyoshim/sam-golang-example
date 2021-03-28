package repo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
)

func PutItem(
	ctx context.Context,
	svc dynamodbiface.DynamoDBAPI,
	in *user.User,
) (err error) {
	item, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return fmt.Errorf("dynamodbattribute.MarshalMap: %+v", err)
	}

	_, err = svc.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(user.TableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("svc.PutItemWithContext: %+v", err)
	}

	return err
}
