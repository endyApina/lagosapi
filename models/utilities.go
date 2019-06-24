package models

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DBConfig hadles database configuration objects
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

//ResponsePackage shows how the rsponse data over the api will look like
type ResponsePackage struct {
	Code   int          `json:"code"`
	Body   ResponseBody `json:"body"`
	Status string       `json:"status"`
}

//ResponseBody is just a struct that is used to wrap some objects to one
type ResponseBody struct {
	Role []*Roles `json:"roles"`
	User User     `json:"user"`
}

//APIData shows the structure the API wants it's data
type APIData struct {
	User User `json:"user"`
	Body User `json:"body"`
}

//DB stores new DB configuration
var DB = new(DBConfig)

//Conn exports Gorm DB pointer
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
	UserID   uint   `gorm:"type:int(10)" json:"user_id"`
	UserName string `gorm:"type:varchar(100)" json:"user_name"`
	Code     int    `gorm:"type:int(10)" json:"code"`
	Role     string `gorm:"type:varchar(100)" json:"role"`
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

//GetUserRoles gets all roles of a user in an array
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

//CheckUserExists functions returns ture or false using email to verify
func CheckUserExists(u User) bool {
	findEmail := Conn.Where("email = ?", u.Email).Find(&u)
	//If email doesn't exist, proceed to check if username exist
	if findEmail != nil && findEmail.Error != nil {
		return false
	}
	return true
}

//CheckUsername checks if username exists as username is Unique
func CheckUsername(u User) bool {
	findUsername := Conn.Where("username = ?", u.Username).Find(&u)
	//If usernmae doesn't exist?
	if findUsername != nil && findUsername.Error != nil {
		return false
	}
	return true
}

//CheckSuperAdmin checks if a particular user posing to be a superadmin is actually true or false
func CheckSuperAdmin(u User) bool {
	var r Roles
	findRole := Conn.Where("user_id = ? AND code = 99", u.ID).Find(&r)
	//If Role doesn't exist?
	if findRole != nil && findRole.Error != nil {
		return false
	}
	return true
}

//CheckSubAdmin checks if a user is a Sub Admin
func CheckSubAdmin(u User) bool {
	var r Roles
	ifSubAdmin := Conn.Where("user_id = ? AND code = 88", u.ID).Find(&r)
	//if user not a sub Admin?
	if ifSubAdmin != nil && ifSubAdmin.Error != nil {
		return false
	}
	return true
}

//DoesSupAdminExist checks if a sup admin exist in the system or not
func DoesSupAdminExist() bool {
	var r Roles
	ifsupadminexists := Conn.Where("code == 99").Find(&r)
	//If role does not exist
	if ifsupadminexists != nil && ifsupadminexists.Error != nil {
		return false
	}
	return true
}

//ResponseObject for roles and users
type ResponseObject struct {
	Role Roles `json:"role"`
	User User  `json:"user"`
}

//GetUserFromRole gets the user information from the role
func GetUserFromRole(roleArray []Roles) interface{} {

	var u User

	var res []*ResponseObject
	for _, role := range roleArray {
		Conn.Where("id = ?", role.ID).Find(&u)
		response := new(ResponseObject)
		response.Role = role
		response.User = u

		res = append(res, response)
	}

	return res
}
