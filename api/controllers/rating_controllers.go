package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nayonacademy/conferenceTracker/api/models"
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateRating(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("create rating")
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	rating := models.Rating{}
	err = json.Unmarshal(body, &rating)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	rating.Prepared()
	err = rating.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	ratingCreated, err := rating.SaveRating(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, rating.ID))
	responses.JSON(w, http.StatusCreated, ratingCreated)
}

func (server *Server) GetRatings(w http.ResponseWriter, r *http.Request){
	rating := models.Rating{}
	ratings, err := rating.FindAllRating(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, ratings)
}

func (server *Server) GetRating(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	rating := models.Rating{}
	ratingGotten, err := rating.FindRatingByID(server.DB,uint64(uid))
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	responses.JSON(w, http.StatusOK, ratingGotten)
}

func (server *Server) UpdateRating(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	rating := models.Rating{}
	err = server.DB.Debug().Model(models.Rating{}).Where("id = ?", pid).Take(&rating).Error
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, errors.New("rating not found"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	ratingUpdate := models.Rating{}
	err = json.Unmarshal(body, &ratingUpdate)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	ratingUpdate.Prepared()
	err = ratingUpdate.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	ratingUpdate.ID = rating.ID
	ratingUpdated, err := ratingUpdate.UpdateRating(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, ratingUpdated)
}

func (server *Server) DeleteRating(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	rating := models.Rating{}
	err = server.DB.Debug().Model(models.Rating{}).Where("id = ?", pid).Take(&rating).Error
	if err != nil{
		responses.ERROR(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}
	_, err = rating.DeleteRating(server.DB, pid)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent,"")
}