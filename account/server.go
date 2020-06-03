package account

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	gokitjwt "github.com/go-kit/kit/auth/jwt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler{
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	r := mux.NewRouter()
	r.Use(CORS)
	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
		))
	r.Methods("GET").Path("/user/{email}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetUser),
		decodeIdReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
		))
	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetUser),
		decodeIdReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		decodeTokenReq,
		encodeResponse,
		))

	r.Methods("POST").Path("/speaker").Handler(httptransport.NewServer(
		//gokitjwt.NewParser(account.JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.CreateCategory),
		endpoints.CreateSpeaker,
		decodeCreateSpeakerRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	r.Methods("GET").Path("/speaker/{id}").Handler(httptransport.NewServer(
		//gokitjwt.NewParser(account.JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.CreateCategory),
		endpoints.GetSpeaker,
		decodeGetSpeakerRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	//return handlers.CORS()(r)
	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Headers:", "*")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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