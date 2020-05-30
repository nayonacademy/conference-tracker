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
	if user.Email == "" || user.Password == ""{
		return RepoErr
	}
	err := r.db.Create(&User{Email: user.Email, Password: user.Password})
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
	//var token string
	err := r.db.First(&user, "email=? and password=?", email, password)
	if err != nil{
		return "", RepoErr
	}
	//token, err = Sign(email, password)
	//if err != nil{
	//	return "", RepoErr
	//}
	return "user", nil
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository{
	return &repo{
		db:    db,
		logger: log.With(logger,"repo","gorm"),
	}
}