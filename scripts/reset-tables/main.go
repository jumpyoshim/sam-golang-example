package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	ddb "gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest/schema"
)

func main() {
	if os.Getenv("DYNAMODB_LOCAL_ENDPOINT") == "" {
		fmt.Println("set DYNAMODB_LOCAL_ENDPOINT=http://localhost:8800")
	}
	svc := dynamodbtest.NewClient(dynamodbtest.NewClientInput{}).(*dynamodb.DynamoDB)

	tables := map[string]*dynamodb.CreateTableInput{
		schema.TableNameUser: schema.User,
	}

	err := ddb.ResetTables(svc, tables)
	if err != nil {
		panic(err)
	}

	fmt.Println("success.")
}
