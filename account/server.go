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

	//Create New user
	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
		))
	// Get a Single User details
	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetUser),
		decodeEmailReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// User Login and get Bearer Token
	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		decodeTokenReq,
		encodeResponse,
		))
	// Create Speaker
	r.Methods("POST").Path("/speaker").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.CreateSpeaker),
		//endpoints.CreateSpeaker,
		decodeCreateSpeakerRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// Get all speaker
	r.Methods("GET").Path("/speaker").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetAllSpeaker),
		//endpoints.CreateSpeaker,
		decodeGetAllSpeakerRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// Get a single Speaker details
	r.Methods("GET").Path("/speaker/{id}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetSpeaker),
		//endpoints.GetSpeaker,
		decodeGetSpeakerRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// Create Category
	r.Methods("POST").Path("/category").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.CreateCategory),
		decodeCategoryReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// Get All Category
	r.Methods("GET").Path("/category").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetCategories),
		decodeCategoriesReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// Get Single Category
	r.Methods("GET").Path("/category/{id}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.GetCategory),
		decodeCategoryReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
	// Update a Single Category
	r.Methods("POST").Path("/category/{name}").Handler(httptransport.NewServer(
		gokitjwt.NewParser(JwtKeyFunc, jwt.SigningMethodHS256, gokitjwt.StandardClaimsFactory)(endpoints.UpdateCategory),
		decodeUpdateReq,
		encodeResponse,
		append(options, httptransport.ServerBefore(gokitjwt.HTTPToContext()))...,
	))
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