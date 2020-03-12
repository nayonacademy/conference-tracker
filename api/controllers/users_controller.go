package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/nayonacademy/conferenceTracker/api/models"
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"io/ioutil"
	"net/http"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	userCreated, err := user.SaveUser(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}
