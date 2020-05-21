package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/nayonacademy/conference-tracker/account"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	_ "github.com/mattn/go-sqlite3"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
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
	fieldKeys := []string{"method","error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace:"my_group",
		Subsystem: "string_service",
		Name:"request_count",
		Help:"Number of requests received",
	},fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace:   "my_group",
		Subsystem:   "string_service",
		Name:        "request_latency_microseconds",
		Help:        "Total duration of requests in microseconds",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace:   "my_group",
		Subsystem:   "string_service",
		Name:        "count_result",
		Help:        "The result of each count method",
	}, []string{})
	level.Info(logger).Log("msg","service started")
	defer level.Info(logger).Log("msg","service ended")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("sqlite3","./nraboy.db")

		if err != nil{
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, email TEXT, password TEXT)")
		statement.Exec()
	}
	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	//srv = account.instrumentingMiddleware{logger, svc}
	//srv = account.instrumentingMiddleware{requestCount, requestLatency, countResult, srv}
	{
		repository := account.NewRepo(db, logger)
		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	endpoints := account.MakeEndpoints(srv)
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit",<-errs)
}
