package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB	*gorm.DB
	Router	*mux.Router
}

func (server *Server) Initialize(){
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string){
	fmt.Println("Listening to port ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}