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

func (server *Server) CreateConference(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("conference create")
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	conference := models.Conference{}
	err = json.Unmarshal(body, &conference)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	conference.Prepare()
	err = conference.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	conferenceCreated, err := conference.SaveConference(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, conference.ID))
	responses.JSON(w, http.StatusCreated, conferenceCreated)
}

func (server *Server) GetConferences(w http.ResponseWriter, r *http.Request){
	conference := models.Conference{}
	conferences, err := conference.FindAllConference(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusCreated, conferences)
}

func (server *Server) GetConference(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	conference := models.Conference{}
	conferenceGotten, err := conference.FindConferenceByID(server.DB, uint64(uid))
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	responses.JSON(w, http.StatusOK, conferenceGotten)
}

func (server *Server) UpdateConference(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10 , 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	conference := models.Conference{}
	err = server.DB.Debug().Model(models.Conference{}).Where("id = ?", pid).Take(&conference).Error
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	// read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	conferenceUpdate := models.Conference{}
	err = json.Unmarshal(body, &conferenceUpdate)
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	conferenceUpdate.Prepare()
	err = conferenceUpdate.Validate()
	if err != nil{
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	conferenceUpdate.ID = conference.ID
	conferenceUpdated, err := conferenceUpdate.UpdateConference(server.DB)
	if err != nil{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("not update conference"))
		return
	}
	responses.JSON(w, http.StatusOK, conferenceUpdated)
}

func (server *Server) DeleteConference(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	conference := models.Conference{}
	err = server.DB.Debug().Model(models.Conference{}).Where("id = ?", pid).Take(&conference).Error
	if err != nil{
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	_, err = conference.DeleteConference(server.DB, pid)
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d",pid))
	responses.JSON(w, http.StatusNoContent,"")
}