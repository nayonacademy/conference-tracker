package account

import "context"

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, email string)(string, error)
	Login(ctx context.Context, email string, password string)(string, error)
}