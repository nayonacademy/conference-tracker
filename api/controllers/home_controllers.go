package controllers

import (
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Welcome to awesome API for Conference Tracker. Feel free to contact with developer si.nayon@gmail.com")
}