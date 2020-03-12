package controllers

import (
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"net/http"
)

func (server *Server) HealthCheck(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "alive")
}