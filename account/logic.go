package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repostory Repository
	logger log.Logger
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger,"method","CreateUser")
	user := User{
		Email:    email,
		Password: password,
	}
	if err := s.repostory.CreateUser(ctx, user); err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create user", user.Email)
	return "success", nil
}

func (s service) GetUser(ctx context.Context, email string) (string, error) {
	logger := log.With(s.logger,"method","GetUser")
	id, err := s.repostory.GetUser(ctx, email)
	if err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("get user", id)
	return id, nil
}

func (s service) Login(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger,"method","Login")
	token, err := s.repostory.Login(ctx, email, password)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("Login",token)
	return token, nil
}

func NewService(rep Repository, logger log.Logger) Service{
	return &service{
		repostory: rep,
		logger:    logger,
	}
}