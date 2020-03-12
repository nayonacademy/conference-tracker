package controllers

import (
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Welcome to awesome API for Conference Tracker")
}