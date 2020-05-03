package sendgrid

import (
	"encoding/json"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)
type SendGrid struct {
	Subject string
	FromEmail string
	FromName string
	ToEmail string
	ToName string
	PlainTextContent string
	HtmlContent string
}
type Response struct {
	StatusCode  int       `json:"statusCode"`
	Headers     map[string][]string  `json:"headers"`
	Body        string    `json:"body"`
}
func (s *SendGrid) EmailSend() (string, error){
	from := mail.NewEmail(s.FromName, s.FromEmail)
	to := mail.NewEmail(s.ToName, s.ToEmail)
	message := mail.NewSingleEmail(from, s.Subject, to, s.PlainTextContent, s.HtmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return "", nil
	}
		//fmt.Println(response.StatusCode)
		//fmt.Println(response.Body)
		//fmt.Println(response.Headers)
	res := &Response{
		StatusCode: response.StatusCode,
		Headers: response.Headers,
		Body: response.Body,
	}
	content, err := json.Marshal(res)
	return string(content), err

}