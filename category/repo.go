package category

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	db *gorm.DB
	logger log.Logger
}

func (r repo) CreateCategory(ctx context.Context, category Category) error {
	var cat Category
	err := r.db.Find(&category, "name = ?", category.Name).Scan(&cat).Error
	if cat.Name == category.Name{
		return errors.New("category already exists")
	}
	//Create
	err = r.db.Create(&Category{Name:category.Name}).Error
	if err != nil{
		return RepoErr
	}
	return nil
}

func (r repo) GetCategory(ctx context.Context, id string) (string, error) {
	var name string
	var category Category
	err := r.db.Find(&category, "id = ?",id)
	if err != nil{
		return "", RepoErr
	}
	return name, nil
}

func (r repo) UpdateCategory(ctx context.Context, name string) (string, error) {
	var category Category
	err := r.db.Where("name = ?", name).Find(&category).Update("name", name).Error
	if err != nil{
		return "", RepoErr
	}
	return category.Name, nil
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository{
	return &repo{
		db: db,
		logger : log.With(logger, "repo","gorm"),
	}
}