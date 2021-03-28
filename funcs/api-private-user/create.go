package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/user/repo"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/funcs"
)

type createOutput struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func create(c *gin.Context) {
	var (
		in  *user.UserInput
		out createOutput
	)

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(500, funcs.Error{Message: err.Error()})
		return
	}

	err = json.Unmarshal(body, &in)
	if err != nil {
		c.JSON(500, funcs.Error{Message: err.Error()})
		return
	}

	u := user.NewUser(in)
	err = repo.PutItem(ctx, fctx.Svc.DynamoDB, u)
	if err != nil {
		c.JSON(500, funcs.Error{Message: err.Error()})
		return
	}

	out = createOutput{
		UUID:      u.UUID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt / 1e9, // NOTE: convert to second
		UpdatedAt: u.UpdatedAt / 1e9, // NOTE: convert to second
	}

	c.JSON(201, out)
}
