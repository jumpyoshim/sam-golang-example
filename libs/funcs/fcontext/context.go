package fcontext

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Context struct {
	Svc Services
}

type Services struct {
	DynamoDB dynamodbiface.DynamoDBAPI
}

type key struct{}

var contextKey = &key{}

func NewContext(parent context.Context, child *Context) context.Context {
	return context.WithValue(parent, contextKey, child)
}

func FromContext(ctx context.Context) (*Context, bool) {
	hcctx, ok := ctx.Value(contextKey).(*Context)
	return hcctx, ok
}

func Must(hcctx *Context, ok bool) *Context {
	if !ok {
		panic("hccontext: can't fetch hcctx")
	}
	return hcctx
}
