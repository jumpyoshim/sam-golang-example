package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/blog"
)

type blogPostInput struct {
	UserUID string `json:"user_uid"`
}

type blogPostOutput struct {
	Item blog.Blog `json:"item"`
}

func blogPost(c *gin.Context) {
	var (
		in  blog.BlogInput
		out blogPostOutput
	)
	c.BindJSON(&in)

	b := blog.NewBlog(&in)

	out.Item = *b

	c.JSON(200, out)
}
