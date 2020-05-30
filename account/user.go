package account

import (
	"context"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, email string) (string, error)
	Login(ctx context.Context, email string, password string)(string, error)
}