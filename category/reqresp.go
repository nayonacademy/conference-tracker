package category

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type(
	CreateCategoryRequest struct {
		Name string `json:"name"`
	}
	CreateCategoryResponse struct {
		OK string	`json:"ok"`
	}
	GetCategoryRequest struct {
		Id string	`json:"id"`
	}
	GetCategoryResponse struct {
		Name string `json:"name"`
	}
	UpdateCategoryRequest struct {
		Name string `json:"name"`
	}
	UpdateCategoryResponse struct {
		Name string `json:"name"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error{
	return json.NewEncoder(w).Encode(response)
}

func decodeCategoryReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req CreateCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
}
func decodeGetReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req GetCategoryRequest
	vars := mux.Vars(r)
	req = GetCategoryRequest{Id:vars["id"]}
	return req, nil
}
func decodeUpdateReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req UpdateCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		return nil, err
	}
	return req, nil
}