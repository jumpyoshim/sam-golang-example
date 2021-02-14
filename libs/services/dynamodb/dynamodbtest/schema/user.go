package schema

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TableNameUser = "user"
)

var User = &dynamodb.CreateTableInput{
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{AttributeName: aws.String("uuid"), AttributeType: aws.String("S")},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{AttributeName: aws.String("uuid"), KeyType: aws.String(dynamodb.KeyTypeHash)},
	},
	BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
}

func init() {
	User.TableName = aws.String(TableNameUser)
}
