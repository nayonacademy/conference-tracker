package category
import (
	"context"
)
type Service interface {
	CreateCategory(ctx context.Context, name string) (string, error)
	//GetCategories(ctx context.Context) (string, error)
	GetCategory(ctx context.Context, id string)(string, error)
	UpdateCategory(ctx context.Context, name string)(string, error)
}