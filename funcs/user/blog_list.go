package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/blog"
)

type blogListInput struct {
	UserUID string `json:"user_uid"`
}

type blogListOutput struct {
	Items []blog.Blog `json:"items"`
}

func blogList(c *gin.Context) {
	var (
		out blogListOutput
	)

	out.Items = []blog.Blog{
		{Title: "Hello, world!", Text: "Hello, world!"},
	}

	c.JSON(200, out)
}
