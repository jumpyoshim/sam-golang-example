package user

import (
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/tweet"
)

type User struct {
	UID    string        `json:"uid" binding:"required,max=64"`
	Email  string        `json:"email" binding:"required,max=256"`
	Name   string        `json:"name" binding:"required,max=128"`
	Tweets []tweet.Tweet `json:"tweets"`
}
