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
	uid := c.Param("uid")

	out.Item = user.User{UID: uid, Email: "foo@example.com"}

	c.JSON(200, out)
}
