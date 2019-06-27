package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"

	//mysql driver

	jwt "github.com/dgrijalva/jwt-go"
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

//APIResponseData stores response data
type APIResponseData struct {
	Code  int          `json:"code"`
	Body  ResponseBody `json:"body"`
	Token string       `json:"token"`
}

//Invite stores json for invitation
type Invite struct {
	Role  int    `json:"role"`
	Email string `json:"email"`
	Token string `json:"token"`
}

//BusinessIdea hold data for Adding an Idea
type BusinessIdea struct {
	gorm.Model
	UserID   int    `gorm:"type:int(10)" json:"user_id" form:"user_id"`
	UserName string `gorm:"type:varchar(100)" json:"user_name" form:"user_name"`
	BusName  string `gorm:"type:varchar(100)" json:"bus_name" form:"bus_name"`
	Email    string `gorm:"type:varchar(100)" json:"email" form:"email"`
	Pitch    string `gorm:"type:varchar(100)" json:"pitch" form:"pitch"`
	Fund     string `gorm:"type:varchar(100)" json:"fund" form:"fund"`
	Industry string `gorm:"type:varchar(100)" json:"industry" form:"industry"`
	Vision   string `gorm:"type:varchar(100)" json:"vision" form:"vision"`
	Mission  string `gorm:"type:varchar(100)" json:"mission" form:"mission"`
	Document string `gorm:"type:varchar(100)" json:"document" form:"document"`
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

	//////////////********** Upload Industry Data *************/////////////////

	var i Industry
	Conn.AutoMigrate(&Industry{})
	checkIndustries := Conn.Find(&i)
	if checkIndustries != nil && checkIndustries.Error != nil {
		industryCSV, err := os.Open("files/industries/industries.csv")
		if err != nil {
			panic(err)
		}

		file := csv.NewReader(industryCSV)
		file.FieldsPerRecord = -1
		rawCSV, err := file.ReadAll()
		if err != nil {
			panic(err)
		}
		var counter uint
		counter = 0

		for _, industry := range rawCSV {
			counter++
			i.ID = counter
			i.Industry = industry[0]
			Conn.Create(&i)
		}

		return
	}

	/////////////END of Industry section /////////////////

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
		Code int    `json:"code"`
		Body string `json:"body"`
	}

	var responseData Message
	responseData.Code = code
	responseData.Body = message

	return responseData
}

//ValidResponse sends valid response to the frontend
func ValidResponse(code int, responseText interface{}) interface{} {
	type responseObject struct {
		Code   int         `json:"code"`
		Body   interface{} `json:"body"`
		Status string      `json:"status"`
		// Role uint
	}

	var sendResponse responseObject
	sendResponse.Code = code
	sendResponse.Body = responseText
	sendResponse.Status = "Success"
	// sendResponse.Role = role

	return sendResponse
}

//APIResponse handles the object to be sent back
func APIResponse(code int, body interface{}, token string) interface{} {
	type response struct {
		Code  int         `json:"code"`
		Body  interface{} `json:"body"`
		Token string      `json:"token"`
	}

	var send response
	send.Code = code
	send.Body = body
	send.Token = token

	return send
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

//ValidUser stores User information with Token string in a single object.
func ValidUser(u interface{}, token string) interface{} {
	type validobject struct {
		Body  interface{} `json:"body"`
		Token string      `json:"token"`
	}

	var r validobject
	r.Body = u
	r.Token = token

	return r
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

//AddUserToken adds a token to a user in a struct
func AddUserToken(u User, token string) interface{} {
	type userToken struct {
		User  User   `json:"user"`
		Token string `json:"string"`
	}

	var uToken userToken
	uToken.User = u
	uToken.Token = token

	return uToken
}

//GetTokenString generates and returns a string.
func GetTokenString(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"secret": username,
		"time":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("jwtkey")))
	if err != nil {
		panic(err)
	}

	return tokenString
}

//ValidateToken validate tokenString
var ValidateToken = func(tokenString string, username string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(beego.AppConfig.String("jwtkey")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["secret"], claims["time"])
	} else {
		log.Println(err)
	}
}

//ValidateRegistration validates the data sent on registration to see if it's valid
func ValidateRegistration(u User) bool {
	if u.FullName == "" {
		return false
	}
	if u.Email == "" {
		return false
	}
	if u.Username == "" {
		return false
	}
	if u.Password == "" {
		return false
	}

	return true
}
