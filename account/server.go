package account

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler{
	key := []byte("secret")
	keys := func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
	jwtOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(gokitjwt.HTTPToContext()),
	}

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
		))
	r.Methods("GET").Path("/user/{email}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(keys, jwt.SigningMethodHS256, gokitjwt.MapClaimsFactory)(endpoints.GetUser),
		decodeIdReq,
		encodeResponse,
		jwtOptions...,
		))
	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		gokitjwt.NewParser(keys, jwt.SigningMethodHS256, gokitjwt.MapClaimsFactory)(endpoints.GetUser),
		decodeIdReq,
		encodeResponse,
		jwtOptions...,
	))
	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		decodeTokenReq,
		encodeResponse,
		))
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}