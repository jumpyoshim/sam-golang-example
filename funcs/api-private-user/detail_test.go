package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user/repo"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs/fcontext"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest/schema"
)

func TestDetail(t *testing.T) {
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

	users := []user.User{
		{
			UUID:      "u01",
			Email:     "foo@example.com",
			Name:      "foo",
			CreatedAt: 1257894000000000000,
			UpdatedAt: 1257894000000000000,
		},
	}
	_, err := repo.BatchWrite(ctx, ddb, users)
	assert.Nil(t, err)

	tests := []struct {
		name       string
		uuid       string
		statusCode int
		want       interface{}
	}{
		{
			name:       "ok",
			uuid:       "u01",
			statusCode: 200,
			want: detailOutput{
				UUID:      "u01",
				Email:     "foo@example.com",
				Name:      "foo",
				CreatedAt: 1257894000,
				UpdatedAt: 1257894000,
			},
		},
		{
			name:       "not_found",
			uuid:       "u02",
			statusCode: 404,
			want: funcs.Error{
				Message: "User does not exist.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := events.APIGatewayProxyRequest{
				Path:       strings.Join([]string{"user", tt.uuid}, "/"),
				HTTPMethod: "GET",
				PathParameters: map[string]string{
					"uuid": tt.uuid,
				},
			}
			got, err := Handler(ctx, req)
			assert.Nil(t, err)
			assert.Equal(t, tt.statusCode, got.StatusCode)

			want, err := json.Marshal(tt.want)
			assert.Equal(t, string(want), got.Body)
		})
	}
}
