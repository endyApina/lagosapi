package models

import (
	"lagosapi/controllers/mailer"
	"log"

	"github.com/astaxie/beego"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//Industry struct holds struct data
type Industry struct {
	gorm.Model
	Industry string `gorm:"type:varchar(100)" json:"industry"`
}

//ValidateToken validates a token string
func ValidateToken(tokenString string) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})

	if err != nil {
		panic(err)
	}

	for key, val := range claims {
		if key == "secret" {
			log.Println(val)
		}
	}

	return
}

//SendRegistrationEmail send confirmation registration email to the registered user.
func SendRegistrationEmail(u User) {
	path := beego.AppConfig.String("templatepath") + "registration.html"

	mailSubject := "Registration to SMEHub Successful"
	newRequest := mailer.NewRequest(u.Email, mailSubject)
	data := mailer.Data{}
	data.User = u.FullName

	go newRequest.Send(path, data)

	return
}

//SendInviteMessage sends an invite message to a user to join the system.
func SendInviteMessage(u User, link string, template string, userType string) {
	mailSubject := "Invitation to Join SMEHUB"
	newRequest := mailer.NewRequest(u.Email, mailSubject)
	data := mailer.Invitation{}
	data.User = u.FullName
	data.Type = userType
	data.Link = link

	go newRequest.Invite(template, data)
	return
}

//GetUserType gets type of a user
func GetUserType(code int) string {
	if code == 999 {
		return "Application Owner"
	}
	if code == 99 {
		return "Super Admin"
	}
	if code == 88 {
		return "Administrative"
	}
	if code == 77 {
		return "Judge"
	}
	if code == 66 {
		return "Mentor"
	}
	if code == 55 {
		return "Investor"
	}
	if code == 0 {
		return "Regular User"
	}

	return "Invalid User Code"
}
