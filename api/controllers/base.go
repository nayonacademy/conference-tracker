package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nayonacademy/conferenceTracker/api/models"
	"log"
	"net/http"
)

type Server struct {
	DB	*gorm.DB
	Router	*mux.Router
}

func (server *Server) Initialize(){
	var err error
	server.DB, err = gorm.Open("sqlite3","./gorm.db")
	if err != nil{
		fmt.Printf("cannot connect to database\n")
		log.Fatal("this is error :",err)
	}else{
		fmt.Printf("we are connected to database\n")
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Category{}, &models.Speaker{})
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string){
	fmt.Println("Listening to port ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}