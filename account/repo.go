package account

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

func (r *repo) CreateUser(ctx context.Context, user User) error {
	var usr User
	err := r.db.First(&user, "email = ?", user.Email).Scan(&usr).Error

	// Create
	err = r.db.Create(&User{Email: user.Email, Password: user.Password}).Error
	if err != nil{
		return RepoErr
	}
	return nil
}

func (r *repo) GetUser(ctx context.Context, email string) (string, error) {
	var id string
	var user User
	err := r.db.First(&user, "email = ?", email)
	if err != nil{
		return "", RepoErr
	}
	return id, nil
}

func (r *repo) Login(ctx context.Context, email string, password string) (string, error) {
	var user User
	var token string
	err := r.db.Where("email = ? and password = ?", email, password).Find(&user).Error
	if err != nil{
		return "", RepoErr
	}
	token , errs := Sign(email, password)
	if errs != nil{
		return "", RepoErr
	}
	return token, nil
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository{
	return &repo{
		db:    db,
		logger: log.With(logger,"repo","gorm"),
	}
}