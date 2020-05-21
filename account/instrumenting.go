package account

import (
"context"
"fmt"
"github.com/go-kit/kit/metrics"
"time"
)

type instrumentingMiddleware struct {
	requestCount metrics.Counter
	requestLatency metrics.Histogram
	countResult metrics.Histogram
	next Service
}

func (mw instrumentingMiddleware) CreateUser(ctx context.Context, email string, password string)(output string, err error){
	defer func(begin time.Time) {
		lvs := []string{"method","createuser","error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.next.CreateUser(ctx, email, password)
	return
}

func (mw instrumentingMiddleware) GetUser(ctx context.Context, email string)(output string, err error){
	defer func(begin time.Time) {
		lvs := []string{"method","getuser","error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.next.GetUser(ctx, email)
	return
}

func (mw instrumentingMiddleware) Login(ctx context.Context, email string, password string)(output string, err error){
	defer func(begin time.Time) {
		lvs := []string{"method","Login","error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.next.Login(ctx, email, password)
	return
}