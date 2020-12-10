package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/tweet"
)

type tweetsInput struct {
	UserUID string `json:"uid"`
}

type tweetsOutput struct {
	Items []tweet.Tweet `json:"items"`
}

func tweets(c *gin.Context) {
	var (
		out tweetsOutput
	)

	out.Items = []tweet.Tweet{
		{Text: "Hello, world!"},
	}

	c.JSON(200, out)
}
