package controllers

import "github.com/nayonacademy/conferenceTracker/api/middlewares"

func (server *Server) InitializeRoutes(){
	server.Router.HandleFunc("/status", middlewares.SetMiddlewareJSON(server.HealthCheck)).Methods("GET")
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")
	// User login and Create users
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")
	server.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	// Category Create , Update and Delete
	server.Router.HandleFunc("/categories", middlewares.SetMiddlewareJSON(server.CreateCategory)).Methods("POST")
	server.Router.HandleFunc("/categories", middlewares.SetMiddlewareJSON(server.GetCategories)).Methods("GET")
	server.Router.HandleFunc("/categories/{id}", middlewares.SetMiddlewareJSON(server.GetCategory)).Methods("GET")
	server.Router.HandleFunc("/categories/{id}", middlewares.SetMiddlewareJSON(server.UpdateCategory)).Methods("PUT")
	server.Router.HandleFunc("/categories/{id}", middlewares.SetMiddlewareJSON(server.DeleteCategory)).Methods("DELETE")
	// Speaker create, update, delete
	server.Router.HandleFunc("/speakers", middlewares.SetMiddlewareJSON(server.CreateSpeaker)).Methods("POST")
	server.Router.HandleFunc("/speakers", middlewares.SetMiddlewareJSON(server.GetSpeakers)).Methods("GET")
	server.Router.HandleFunc("/speakers/{id}", middlewares.SetMiddlewareJSON(server.GetSpeaker)).Methods("GET")
	server.Router.HandleFunc("/speakers/{id}", middlewares.SetMiddlewareJSON(server.UpdateSpeaker)).Methods("PUT")
	server.Router.HandleFunc("/speakers/{id}", middlewares.SetMiddlewareJSON(server.DeleteSpeaker)).Methods("DELETE")
	// Conference create, update, delete
	server.Router.HandleFunc("/conferences", middlewares.SetMiddlewareJSON(server.CreateConference)).Methods("POST")
	server.Router.HandleFunc("/conferences", middlewares.SetMiddlewareJSON(server.GetConferences)).Methods("GET")
	server.Router.HandleFunc("/conferences/{id}", middlewares.SetMiddlewareJSON(server.GetConference)).Methods("GET")
	server.Router.HandleFunc("/conferences/{id}", middlewares.SetMiddlewareJSON(server.UpdateConference)).Methods("PUT")
	server.Router.HandleFunc("/conferences/{id}", middlewares.SetMiddlewareJSON(server.DeleteConference)).Methods("DELETE")
	// verification
	server.Router.HandleFunc("/verification/{key}", middlewares.SetMiddlewareJSON(server.Verification)).Methods("GET")
	// send grid email send test
	server.Router.HandleFunc("/email", middlewares.SetMiddlewareJSON(server.EmailSend)).Methods("POST")
	// All test and RND things
	server.Router.HandleFunc("/mytest", middlewares.SetMiddlewareJSON(server.MyTest)).Methods("GET")
}
