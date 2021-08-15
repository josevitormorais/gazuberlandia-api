package gazuberlandia

import (
	"context"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) error
	FindUserById(ctx context.Context, id int) ([]User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindAllUsers(ctx context.Context) ([]User, error)
}
