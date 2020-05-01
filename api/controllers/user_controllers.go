package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nayonacademy/conferenceTracker/api/models"
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"github.com/nayonacademy/conferenceTracker/api/sendgrid"
	"io/ioutil"
	"net/http"
	"strconv"
	"crypto/rand"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, errors.New("format error"))
		return
	}
	// verification key generate and send in email and store in redisservice
	verification_key := make([]byte, 14)
	_, err = rand.Read(verification_key)
	if err != nil {
		return
	}
	email_send := sendgrid.SendGrid{
		Subject:          "Account verification",
		FromEmail:        "info@conferencetracker.com",
		FromName:         "Admin",
		ToEmail:          userCreated.Email,
		ToName:           userCreated.Nickname,
		PlainTextContent: "",
		HtmlContent:      string(verification_key),
	}
	_, _ = email_send.EmailSend()

	//email_send.
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}