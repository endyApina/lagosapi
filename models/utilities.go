package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ResponsePackage struct {
	Code   int          `json:"code"`
	Body   ResponseBody `json:"body"`
	Status string       `json:"status"`
}

type ResponseBody struct {
	Role []*Roles `json:"roles"`
	User User     `json:"user"`
}

type ApiData struct {
	User User `json:"user"`
	Body User `json:"body"`
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

//Roles struct handles database representation of the model.
type Roles struct {
	gorm.Model
	UserID   uint   `gorm:"type:int(10)"`
	UserName string `gorm:"type:varchar(100)"`
	Code     int    `gorm:"type:int(10)"`
	Role     string `gorm:"type:varchar(100)"`
}

//Response sends bad response data to you
func Response(code int, message string) interface{} {
	type Message struct {
		Code int
		Body string
	}

	var responseData Message
	responseData.Code = code
	responseData.Body = message

	return responseData
}

//ValidResponse sends valid response to the frontend
func ValidResponse(code int, responseText interface{}) interface{} {
	type responseObject struct {
		Code   int
		Body   interface{}
		Status string
		// Role uint
	}

	var sendResponse responseObject
	sendResponse.Code = code
	sendResponse.Body = responseText
	sendResponse.Status = "Success"
	// sendResponse.Role = role

	return sendResponse
}

//CreateDefaultRole create default role of 00 for regular users
func CreateDefaultRole(u User) Roles {
	var role Roles
	role.UserID = u.ID
	role.UserName = u.FullName
	role.Code = u.Role

	return role
}

//AssociateRoleUser merges users with thier roles
func AssociateRoleUser(r interface{}, u User) interface{} {
	type response struct {
		Role interface{}
		User interface{}
	}

	var respond response
	respond.Role = r
	respond.User = u

	return respond
}

func GetUserRoles(u User) interface{} {
	type response struct {
		Role []*Roles
		User interface{}
	}

	var res response
	res.User = u

	Conn.Where("user_id = ?", u.ID).Find(&res.Role)
	return res
}

func CheckUserExists(u User) bool {
	findEmail := Conn.Where("email = ?", u.Email).Find(&u)
	//If email doesn't exist, proceed to check if username exist
	if findEmail != nil && findEmail.Error != nil {
		return false
	} else {
		return true
	}
}

func CheckUsername(u User) bool {
	findUsername := Conn.Where("username = ?", u.Username).Find(&u)
	//If usernmae doesn't exist?
	if findUsername != nil && findUsername.Error != nil {
		return false
	} else {
		return true
	}
}

func CheckSuperAdmin(u User) bool {
	var r Roles
	findRole := Conn.Where("user_id = ? AND code = 99", u.ID).Find(&r)
	//If Role doesn't exist?
	if findRole != nil && findRole.Error != nil {
		return false
	} else {
		return true
	}
}

func CheckSubAdmin(u User) bool {
	var r Roles
	ifSubAdmin := Conn.Where("user_id = ? AND code = 88", u.ID).Find(&r)
	//if user not a sub Admin?
	if ifSubAdmin != nil && ifSubAdmin.Error != nil {
		return false
	} else {
		return true
	}
}
