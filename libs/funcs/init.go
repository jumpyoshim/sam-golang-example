package funcs

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs/fcontext"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/services/dynamodb/dynamodbtest"
)

func Init(ctx context.Context) context.Context {
	if !IsLocal() {
		ctx = initInServices(ctx)
	} else {
		ctx = initInLocal(ctx)
	}

	return ctx
}

func IsLocal() (ok bool) {
	env := os.Getenv("APP_ENV")
	if env == "local" {
		ok = true
	}

	return ok
}

func initInServices(ctx context.Context) context.Context {
	sess := session.Must(session.NewSession())

	svcd := dynamodb.New(sess)

	fctx := &fcontext.Context{
		Svc: fcontext.Services{
			DynamoDB: svcd,
		},
	}
	ctx = fcontext.NewContext(ctx, fctx)

	return ctx
}

func initInLocal(ctx context.Context) context.Context {
	ddb := dynamodbtest.NewClient(dynamodbtest.NewClientInput{}).(*dynamodb.DynamoDB)

	ctx = fcontext.NewContext(
		context.Background(),
		&fcontext.Context{
			Svc: fcontext.Services{DynamoDB: ddb},
		},
	)

	return ctx
}
