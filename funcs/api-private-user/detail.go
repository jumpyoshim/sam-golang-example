package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user/repo"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs"
)

type detailOutput struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func detail(c *gin.Context) {
	var (
		u   *user.User
		out detailOutput
	)

	uuid := c.Param("uuid")
	key := user.UserUUIDKey{UUID: uuid}

	u, err := repo.GetItem(ctx, fctx.Svc.DynamoDB, key)
	if err != nil {
		c.JSON(500, funcs.Error{Message: err.Error()})
		return
	}
	if u == nil {
		c.JSON(404, funcs.Error{Message: "User does not exist."})
		return
	}

	out = detailOutput{
		UUID:      uuid,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt / 1e9, // NOTE: convert to second
		UpdatedAt: u.UpdatedAt / 1e9, // NOTE: convert to second
	}

	c.JSON(200, out)
}
