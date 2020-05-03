package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nayonacademy/conferenceTracker/api/controllers"
	"log"
)

var server = controllers.Server{}

func Run(){
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	server.Initialize()
	server.InitializeRedis()
	server.Run(":8000")
}