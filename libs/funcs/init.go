package funcs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs/fcontext"
)

func Init(ctx context.Context) context.Context {
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
