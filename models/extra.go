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
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})

	if err != nil {
		panic(err)
		log.Println(token)
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
