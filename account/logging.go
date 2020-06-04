package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next Service
}

func (mw loggingMiddleware) CreateUser(ctx context.Context, email string, password string)(output string, err error){
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method","createuser",
			"input", email, password,
			"output",output,
			"err", err,
			"took", time.Since(begin),
			)
	}(time.Now())
	output, err = mw.next.CreateUser(ctx, email, password)
	return
}

func (mw loggingMiddleware) GetUser(ctx context.Context, id string)(output User, err error){
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method","createuser",
			"input", id,
			"output",output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.GetUser(ctx, id)
	return
}

func (mw loggingMiddleware) Login(ctx context.Context, email string, password string)(output string, err error){
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method","login",
			"input", email, password,
			"output",output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Login(ctx, email, password)
	return
}