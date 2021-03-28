package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs/fcontext"
)

var (
	ginLambda *ginadapter.GinLambda
	ctx       context.Context
	fctx      *fcontext.Context
)

func init() {
	r := gin.Default()
	r.POST("/user", create)
	r.GET("/user/:uuid", detail)

	ginLambda = ginadapter.New(r)
}

func Handler(c context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx = funcs.Init(c)
	fctx = fcontext.Must(fcontext.FromContext(ctx))

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
