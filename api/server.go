package api

import "github.com/nayonacademy/conferenceTracker/api/controllers"

var server = controllers.Server{}

func Run(){
	server.Initialize()
	server.Run(":8000")
}