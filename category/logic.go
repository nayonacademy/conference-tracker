package category

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repostory Repository
	logger log.Logger
}

func (s service) GetCategory(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger,"method","GetCategory")
	name, err := s.repostory.GetCategory(ctx, id)
	if err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("get category", name)
	return name, nil
}

func (s service) UpdateCategory(ctx context.Context, name string) (string, error) {
	logger := log.With(s.logger, "method","UpdateCategory")
	name, err := s.repostory.UpdateCategory(ctx, name)
	if err != nil{
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("update category", name)
	return name, nil
}

func (s service) CreateCategory(ctx context.Context, name string)(string, error){
	logger := log.With(s.logger,"method","CreateCategory")
	category := Category{
		Name:  name,
	}
	if err := s.repostory.CreateCategory(ctx, category); err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create user", category.Name)
	return "success", nil
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}