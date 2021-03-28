package dynamodbtest

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type ResetTablesInput struct {
	Tables map[string]*dynamodb.CreateTableInput
}

func ResetTables(svc dynamodbiface.DynamoDBAPI, tables map[string]*dynamodb.CreateTableInput) (err error) {
	for _, table := range tables {
		_, err = svc.DeleteTable(&dynamodb.DeleteTableInput{TableName: table.TableName})
		if err != nil {
			if _, ok := err.(*dynamodb.ResourceNotFoundException); !ok {
				return fmt.Errorf("svc.DeleteTable: %+v", err)
			}
		}

		_, err = svc.CreateTable(table)
		if err != nil {
			return fmt.Errorf("svc.CreateTable: %+v", err)
		}
	}

	return err
}
