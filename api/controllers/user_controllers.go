package controllers

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nayonacademy/conferenceTracker/api/models"
	"github.com/nayonacademy/conferenceTracker/api/responses"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"

	"io/ioutil"

	"net/http"
	"os"
	"strconv"
	"time"
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
	emailCheck := user.EmailExists(server.DB,user.Email)
	if emailCheck == true {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("email already exists"))
		return
	}
	nickName := user.UserExists(server.DB,user.Nickname)
	if nickName == true {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("nickname already exists"))
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		responses.ERROR(w, http.StatusInternalServerError, errors.New("format error"))
		return
	}
	//verification key generate and send in email and store in redis Service
	verification_key := make([]byte, 14)
	_, err = rand.Read(verification_key)
	if err != nil {
		return
	}
	//email_send := sendgrid.SendGrid{
	//	Subject:          "Account verification",
	//	FromEmail:        "info@conferencetracker.com",
	//	FromName:         "Admin",
	//	ToEmail:          userCreated.Email,
	//	ToName:           userCreated.Nickname,
	//	PlainTextContent: "",
	//	HtmlContent:      string(verification_key),
	//}
	//_, _ = email_send.EmailSend()
	//redis write the secret key
	key1 := string(verification_key)
	fmt.Println(key1)
	value1 := &RedisValue{Name: userCreated.Nickname, Email: userCreated.Email}
	err = server.setKey(key1, value1, time.Minute*1)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
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

func (server *Server) Verification(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key1 := vars["key"]
	key1 = "sampleKey"
	//value1 := &RedisValue{Name: "someName", Email: "someemail@abc.com"}
	//err := server.setKey(key1, value1, time.Minute*1)
	//if err != nil {
	//	log.Fatalf("Error: %v", err.Error())
	//}
	value2 := &RedisValue{}
	err := server.getKey(key1, value2)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	//email := ""
	//user := models.User{}
	//userGotten, err := user.FindUserByEmail(server.DB, email)
	//if err != nil {
	//	responses.ERROR(w, http.StatusBadRequest, err)
	//	return
	//}
	responses.JSON(w, http.StatusOK, value2)
}

func (server *Server) EmailSend(w http.ResponseWriter, r *http.Request){
	//myKey := make([]byte, 14)
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "si.nayon@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	fmt.Println("api_key", os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))
	responses.JSON(w, http.StatusOK,  "Email send successfully")
}