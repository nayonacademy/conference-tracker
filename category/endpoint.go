package category

import (
	"github.com/go-kit/kit/endpoint"
	"context"
)

type Endpoints struct {
	CreateCategory endpoint.Endpoint
	GetCategory endpoint.Endpoint
	UpdateCategory endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints{
	return Endpoints{
		CreateCategory: makeCreateCategoryEndpoints(s),
		GetCategory:    makeGetCategoryEndpoints(s),
		UpdateCategory: makeUpdateCategoryEndpoints(s),
	}
}

func makeCreateCategoryEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateCategoryRequest)
		ok, err := s.CreateCategory(ctx, req.Name)
		return CreateCategoryResponse{OK:ok}, err
	}
}

func makeGetCategoryEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(response interface{}, err error){
		req := request.(GetCategoryRequest)
		name, err := s.GetCategory(ctx, req.Id)
		return GetCategoryResponse{Name:name}, err
	}
}

func makeUpdateCategoryEndpoints(s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateCategoryRequest)
		cat, err := s.UpdateCategory(ctx, req.Name)
		return UpdateCategoryResponse{Name:cat}, err
	}
}