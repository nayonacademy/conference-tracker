package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/log"
	"time"
)

var RepoErr = errors.New("unable to handle repo request")

type repo struct {
	db *sql.DB
	logger log.Logger
}

func (r *repo) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users (id, email, password)
			VALUES($1,$2,$3)`
	if user.Email == "" || user.Password == ""{
		return RepoErr
	}
	_, err := r.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil{
		return err
	}
	return nil
}

func (r *repo) GetUser(ctx context.Context, email string) (string, error) {
	var id string
	err := r.db.QueryRow("SELECT id FROM users WHERE email=$1",email).Scan(&id)
	if err != nil{
		return "", RepoErr
	}
	return id, nil
}

func (r *repo) Login(ctx context.Context, email string, password string) (string, error) {
	fmt.Println("repo: step 3", email, password)
	var id string
	err := r.db.QueryRow("SELECT id FROM users WHERE email=$1 and password=$2",email,password).Scan(&id)
	if err != nil{
		return "", RepoErr
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil{
		return "", RepoErr
	}
	return t, nil
}

func NewRepo(db *sql.DB, logger log.Logger) Repository{
	return &repo{
		db:    db,
		logger: log.With(logger,"repo","sql"),
	}
}