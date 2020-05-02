package controllers

import (
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"net/http"
)

func (server *Server) HealthCheck(w http.ResponseWriter, r *http.Request){
	//key1 := "sampleKey"
	//value1 := &RedisValue{Name: "someName", Email: "nayon@abc.com"}
	//err := server.setKey(key1, value1, time.Minute*1)
	//if err != nil {
	//	log.Fatalf("Error: %v", err.Error())
	//}
	//value2 := &RedisValue{}
	//err = server.getKey(key1, value2)
	//if err != nil {
	//	log.Fatalf("Error: %v", err.Error())
	//}
	responses.JSON(w, http.StatusOK, "alive")
}