package dynamodbtest

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest/schema"
)

type SetupInput struct {
	Svc    *dynamodb.DynamoDB
	Tables map[string]*dynamodb.CreateTableInput
}

func Setup(in SetupInput) (svc *dynamodb.DynamoDB) {
	var (
		tables = in.Tables
	)

	svc = in.Svc
	if svc == nil {
		svc = NewClient(NewClientInput{}).(*dynamodb.DynamoDB)
	}

	if tables == nil {
		tables = map[string]*dynamodb.CreateTableInput{
			*schema.User.TableName: schema.User,
		}
	}

	err := ResetTables(svc, tables)
	if err != nil {
		panic(err)
	}

	return
}