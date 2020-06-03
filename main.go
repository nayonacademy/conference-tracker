package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nayonacademy/conference-tracker/account"
	"github.com/nayonacademy/conference-tracker/category"
	"github.com/nayonacademy/conference-tracker/conference"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	var httpAddr = flag.String("http",":8080","http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time",log.DefaultTimestampUTC,"caller",log.DefaultCaller,)
	}
	// prometheus implement
	//fieldKeys := []string{"method","error"}
	//requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
	//	Namespace:"my_group",
	//	Subsystem: "string_service",
	//	Name:"request_count",
	//	Help:"Number of requests received",
	//},fieldKeys)
	//requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	//	Namespace:   "my_group",
	//	Subsystem:   "string_service",
	//	Name:        "request_latency_microseconds",
	//	Help:        "Total duration of requests in microseconds",
	//}, fieldKeys)
	//countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	//	Namespace:   "my_group",
	//	Subsystem:   "string_service",
	//	Name:        "count_result",
	//	Help:        "The result of each count method",
	//}, []string{})
	level.Info(logger).Log("msg","service started")
	defer level.Info(logger).Log("msg","service ended")


	var err error
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	defer db.Close()
	db.AutoMigrate(&account.User{}, &category.Category{},&conference.ConfOwnProfile{},&conference.Rating{},&conference.Conference{},&conference.Location{},&conference.Report{},&conference.Speaker{})

	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	var srv_cat category.Service
	var srv_conf conference.Service
	//srv = account.instrumentingMiddleware{logger, svc}
	//srv = account.instrumentingMiddleware{requestCount, requestLatency, countResult, srv}
	{
		repository := account.NewRepo(db, logger)
		cat_repo := category.NewRepo(db, logger)
		conf_repo := conference.NewRepo(db, logger)
		srv = account.NewService(repository, logger)
		srv_cat = category.NewService(cat_repo,logger)
		srv_conf = conference.NewService(conf_repo, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	endpoints := account.MakeEndpoints(srv)
	cat_endpoints := category.MakeEndpoints(srv_cat)
	conf_endpoints := conference.MakeEndpoints(srv_conf)
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		handler = category.NewHTTPServer(ctx, cat_endpoints)
		handler = conference.NewHTTPServer(ctx, conf_endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit",<-errs)
}
