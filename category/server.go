package category

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nayonacademy/conference-tracker/account"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler{
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	r := mux.NewRouter()
	r.Use(CORS)
	r.Methods("POST").Path("/category").Handler(httptransport.NewServer(
		//gokitjwt.NewParser(account.JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.CreateCategory),
		endpoints.CreateCategory,
		decodeCategoryReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
		))
	r.Methods("GET").Path("/category/{id}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(account.JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetCategory),
		decodeGetReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	r.Methods("POST").Path("/category/{name}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(account.JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.UpdateCategory),
		decodeUpdateReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	return r
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")
		// Next
		next.ServeHTTP(w, r)
		return
	})
}