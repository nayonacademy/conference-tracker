package category

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Repository interface {
	CreateCategory(ctx context.Context, category Category) error
	//GetCategories(ctx context.Context) (string, error)
	GetCategory(ctx context.Context, name string)(string, error)
	UpdateCategory(ctx context.Context, name string)(string, error)
}