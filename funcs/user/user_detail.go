package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
)

type userDetailInput struct {
	UserUID string `json:"user_uid"`
}

type userDetailOutput struct {
	Item user.User `json:"item"`
}

func userDetail(c *gin.Context) {
	var (
		out userDetailOutput
	)

	out.Item = user.User{UID: "foo", Email: "foo@example.com"}

	c.JSON(200, out)
}
