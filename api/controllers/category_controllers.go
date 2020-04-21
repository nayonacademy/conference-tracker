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

func(server *Server) CreateCategory(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Create Category")
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	category := models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	category.Prepare()
	err = category.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	categoryCreated, err := category.SaveCategory(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, category.ID))
	responses.JSON(w, http.StatusCreated, categoryCreated)
}

func (server *Server) GetCategories(w http.ResponseWriter, r *http.Request){
	category := models.Category{}

	categories, err := category.FindAllCategory(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, categories)
}

func (server *Server) GetCategory(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	category := models.Category{}
	categoryGotten, err := category.FindCategoryByID(server.DB, uint64(uid))
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	responses.JSON(w, http.StatusOK, categoryGotten)
}

func (server *Server) UpdateCategory(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	// check if the category id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// check if the Category exist
	category := models.Category{}
	err = server.DB.Debug().Model(models.Category{}).Where("id = ?", pid).Take(&category).Error
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, errors.New("category not found"))
		return
	}

	// read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// start processing the request data
	categoryUpdate := models.Category{}
	err = json.Unmarshal(body, &categoryUpdate)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	categoryUpdate.Prepare()
	err = categoryUpdate.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	categoryUpdate.ID = category.ID
	categoryUpdated, err := categoryUpdate.UpdateCategory(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("not updated category"))
		return
	}
	responses.JSON(w, http.StatusOK, categoryUpdated)
}

func (server *Server) DeleteCategory(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	// is a valid category id given to us ?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// check if the category exist
	category := models.Category{}
	err = server.DB.Debug().Model(models.Category{}).Where("id = ?", pid).Take(&category).Error
	if err != nil{
		responses.ERROR(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}
	_, err = category.DeleteACategory(server.DB, pid)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent,"")
}