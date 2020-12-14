package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/blog"
)

type blogsInput struct {
	UserUID string `json:"uid"`
}

type blogsOutput struct {
	Items []blog.Blog `json:"items"`
}

func blogs(c *gin.Context) {
	var (
		out blogsOutput
	)

	out.Items = []blog.Blog{
		{Title: "Hello, world!", Text: "Hello, world!"},
	}

	c.JSON(200, out)
}
