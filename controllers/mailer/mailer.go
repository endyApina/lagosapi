package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	gomail "gopkg.in/gomail.v2"
)

// var auth = smtp.PlainAuth("", "info@my-gpi.com", "oW[Rmakd1rTB", "mail.my-gpi.com")

// Data for composing mail data that will be made available in the mail template
type Data struct {
	Items     interface{}
	User      string
	Password  string
	CreatedAt time.Time
}

//Invest for getting investor details
type Invest struct {
	Name string
}

// Request a request object model
type Request struct {
	to      string
	subject string
	body    string
	//attachment []string
}

// NewRequest for creating new Request object
func NewRequest(to string, subject string /*attachment []string*/) *Request {
	return &Request{
		to:      to,
		subject: subject,
		//attachment: attachment,
	}
}

// sendEmail for setting up email parameters
func (r *Request) sendEmail() bool {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "endy.apina@my-gpi.io", "Global Performance Index")
	m.SetHeader("To", r.to)
	m.SetHeader("Subject", r.subject)
	m.SetBody("text/html", r.body)
	// for _, v := range r.attachment {
	// 	m.Attach(v)
	// }

	d := gomail.NewDialer("mail.my-gpi.io", 587, "endy.apina@my-gpi.io", "newAB0000")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Send for sending out email
func (r *Request) Send(tempName string, item Data) {
	err := r.ParseTemplate(tempName, item)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendEmail(); ok {

	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
		panic(err)
	}
}

// ParseTemplate for parsing email template
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
