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
		{AttributeName: aws.String("email"), AttributeType: aws.String("S")},
		{AttributeName: aws.String("uuid"), AttributeType: aws.String("S")},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{AttributeName: aws.String("email"), KeyType: aws.String(dynamodb.KeyTypeHash)},
	},
	GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
		{
			IndexName: aws.String("user_uuid"),
			KeySchema: []*dynamodb.KeySchemaElement{
				{AttributeName: aws.String("uuid"), KeyType: aws.String(dynamodb.KeyTypeHash)},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
		},
	},
	BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
}

func init() {
	User.TableName = aws.String(TableNameUser)
}
