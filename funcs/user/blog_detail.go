package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/blog"
)

type blogDetailInput struct {
	UserUID string `json:"user_uid"`
}

type blogDetailOutput struct {
	Item blog.Blog `json:"item"`
}

func blogDetail(c *gin.Context) {
	var (
		out blogDetailOutput
	)

	out.Item = blog.Blog{Title: "Hello, world!", Text: "Hello, world!"}

	c.JSON(200, out)
}
