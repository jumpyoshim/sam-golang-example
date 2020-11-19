package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Output struct {
	Message string `json:"message"`
}

func Handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest, // nolint: gocritic
) (out events.APIGatewayProxyResponse, err error) {
	o := Output{
		Message: "ok",
	}

	b, err := json.Marshal(o)
	if err != nil {
		return out, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
