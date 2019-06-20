package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ResponsePackage struct {
	Code   int
	Body   User
	Status string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var DB = new(DBConfig)

var Conn *gorm.DB

func init() {
	DB.Host = "127.0.0.1"
	DB.User = "ssd"
	DB.Password = "ssd"
	DB.Database = "smehub"

	conn, err := gorm.Open("mysql", DB.User+":"+DB.Password+"@/"+DB.Database+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	Conn = conn

}

func Response(code int, message string) interface{} {
	type Message struct {
		Code    int
		Message string
	}

	var responseData Message
	responseData.Code = code
	responseData.Message = message

	return responseData
}

func ValidResponse(code int, responseText interface{}) interface{} {
	type responseObject struct {
		Code   int
		Body   interface{}
		Status string
	}

	var sendResponse responseObject
	sendResponse.Code = code
	sendResponse.Body = responseText
	sendResponse.Status = "Success"

	return sendResponse
}
