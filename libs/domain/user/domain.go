package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	TableName = "user"
)

type User struct {
	UUID      string `json:"uuid" binding:"required"`
	Email     string `json:"email" binding:"required,max=256"`
	Name      string `json:"name" binding:"required,max=32"`
	CreatedAt int64  `json:"created_at" binding:"required"`
	UpdatedAt int64  `json:"updated_at" binding:"required"`
}

type UserEmailKey struct {
	Email string `json:"email"`
}

type UserUUIDKey struct {
	UUID string `json:"uuid"`
}

type UserInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUser(in *UserInput) *User {
	u := &User{
		Email: in.Email,
		Name:  in.Name,
	}

	u.UUID = uuid.New().String()

	now := time.Now()
	u.CreatedAt = now.UnixNano()
	u.UpdatedAt = now.UnixNano()

	return u
}
