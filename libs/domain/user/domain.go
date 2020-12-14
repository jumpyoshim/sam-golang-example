package user

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/jumpyoshim/sam-goloang-example/libs/domain/blog"
)

type User struct {
	UUID      string      `json:"uuid" binding:"required"`
	UID       string      `json:"uid" binding:"required,max=64"`
	Email     string      `json:"email" binding:"required,max=256"`
	Name      string      `json:"name" binding:"required,max=128"`
	Blogs     []blog.Blog `json:"blogs" binding:"max=10000"`
	CreatedAt int64       `json:"created_at" binding:"required"`
	UpdatedAt int64       `json:"updated_at" binding:"required"`
}

type UserInput struct {
	UID   string
	Email string
	Name  string
}

func NewUser(in *UserInput) *User {
	u := &User{
		UID:   in.UID,
		Email: in.Email,
		Name:  in.Name,
	}

	u.UUID = uuid.New().String()

	now := time.Now()
	u.CreatedAt = now.UnixNano()
	u.UpdatedAt = now.UnixNano()

	return u
}
