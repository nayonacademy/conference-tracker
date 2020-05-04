package controllers

import (
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"math/rand"
	"net/http"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func (server *Server) MyTest(w http.ResponseWriter, r *http.Request){
	secret_key := RandStringBytes(25)
	responses.JSON(w, http.StatusOK, "hello I am here"+secret_key)
}
