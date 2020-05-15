package account

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type(
	CreateUserRequest struct {
		Email string `json:"email"`
		Password string	`json:"password"`
	}
	CreateUserResponse struct {
		Ok string	`json:"ok"`
	}
	GetUserRequest struct {
		Email string	`json:"email"`
	}
	GetUserResponse struct {
		Id string	`json:"id"`
	}
	LoginRequest struct {
		Email string	`json:"email"`
		Password string	`json:"password"`
	}
	LoginResponse struct {
		Token string	`json:"token"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error{
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request)(interface{}, error){
 var req CreateUserRequest
 err := json.NewDecoder(r.Body).Decode(&req)
 if err != nil{
 	return nil, err
 }
 return req, nil
}
func decodeIdReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req GetUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
}
func decodeTokenReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req LoginRequest
	vars := mux.Vars(r)
	req = LoginRequest{
		Email:    vars["email"],
		Password: vars["password"],
	}
	return req, nil
}