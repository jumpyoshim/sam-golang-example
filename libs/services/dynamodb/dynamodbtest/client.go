package dynamodbtest

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type NewClientInput struct {
	Sess     *session.Session
	Endpoint string
}

func NewClient(in NewClientInput) dynamodbiface.DynamoDBAPI {
	var (
		sess     = in.Sess
		endpoint = in.Endpoint
	)

	if sess == nil {
		sess = session.Must(session.NewSession(
			&aws.Config{Region: aws.String("ap-northeast-1")},
		))
	}

	if endpoint == "" {
		endpoint = os.Getenv("DYNAMODB_LOCAL_ENDPOINT")
	}
	if endpoint == "" {
		endpoint = "http://dynamodb-local:8000"
	}
	sess.Config.WithEndpoint(endpoint)

	return dynamodb.New(sess)
}
