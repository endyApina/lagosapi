package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/astaxie/beego"

	gomail "gopkg.in/gomail.v2"
)

// Data for composing mail data that will be made available in the mail template
type Data struct {
	Items     interface{}
	User      string
	Password  string
	CreatedAt time.Time
}

//Invite holds data for invitation
type Invitation struct {
	User string
	Type string
	Link string
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
	m.SetAddressHeader("From", beego.AppConfig.String("maileremail"), beego.AppConfig.String("mailerheader"))
	m.SetHeader("To", r.to)
	m.SetHeader("Subject", r.subject)
	m.SetBody("text/html", r.body)
	// for _, v := range r.attachment {
	// 	m.Attach(v)
	// }

	d := gomail.NewDialer(beego.AppConfig.String("mailersmtp"), 587, beego.AppConfig.String("maileremail"), beego.AppConfig.String("mailerpassword"))
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

// Invite for sending out email
func (r *Request) Invite(tempName string, item Invitation) {
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
