package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs/fcontext"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest/schema"
)

func TestCreate(t *testing.T) {
	ddb := dynamodbtest.Setup(dynamodbtest.SetupInput{
		Tables: map[string]*dynamodb.CreateTableInput{
			schema.TableNameUser: schema.User,
		},
	})
	ctx := fcontext.NewContext(
		context.Background(),
		&fcontext.Context{
			Svc: fcontext.Services{DynamoDB: ddb},
		},
	)

	tests := []struct {
		name       string
		in         user.UserInput
		statusCode int
		want       interface{}
	}{
		{
			name: "created",
			in: user.UserInput{
				Email: "foo@example.com",
				Name:  "foo",
			},
			statusCode: 201,
			want: createOutput{
				Email: "foo@example.com",
				Name:  "foo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.in)
			assert.Nil(t, err)

			req := events.APIGatewayProxyRequest{
				Path:       "/user",
				HTTPMethod: "POST",
				Body:       string(b),
			}
			got, err := Handler(ctx, req)
			assert.Nil(t, err)
			assert.Equal(t, tt.statusCode, got.StatusCode)

			var out createOutput
			err = json.Unmarshal([]byte(got.Body), &out)
			assert.Nil(t, err)

			out.UUID = ""
			out.CreatedAt = 0
			out.UpdatedAt = 0
			assert.Equal(t, tt.want, out)
		})
	}
}
