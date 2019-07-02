package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	//UserList returns an array of Users struct
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	// u := User{"Endy Apinageri", "astaxie", "pappi", "male", "03-05-1997", "Primewares", "Lekki", "Nigeria", "astaxie@gmail.com"}
	// UserList["user_11111"] = &u
}

//LoginDetails stores login details
type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//User struct shows models for users
type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(100)" json:"full_name"`
	Username string `gorm:"type:varchar(100) unique" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Gender   string `gorm:"type:varchar(100)" json:"gender"`
	DOB      string `gorm:"type:varchar(100)" json:"dob"`
	Street   string `gorm:"type:varchar(100)" json:"street"`
	City     string `gorm:"type:varchar(100)" json:"city"`
	Country  string `gorm:"type:varchar(100)" json:"country"`
	Email    string `gorm:"type:varchar(100); unique_index" json:"email"`
	Role     int    `gorm:"type:int(10)" json:"role"`
}

//AddUser creates a new user.
func AddUser(u User) interface{} {
	// u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Conn.AutoMigrate(&User{})

	findEmail := Conn.Where("email = ?", u.Email).Find(&u)

	//If email doesn't exist, proceed to check if username exist
	if findEmail != nil && findEmail.Error != nil {
		checkUsername := Conn.Where("username = ?", u.Username).Find(&u)
		if checkUsername != nil && checkUsername.Error != nil {
			//If Email and Username doesn't exist.
			if u.Role != 00 {
				responseData := Response(403, "Unauthorized Role")

				return responseData
			}
			Conn.Create(&u)

			Conn.AutoMigrate(&Roles{})
			getDefaultRole := CreateDefaultRole(u)
			Conn.Create(&getDefaultRole)
			SendRegistrationEmail(u)
			tokenString := GetTokenString(u.Username)
			getRoles := AssociateRoleUser(getDefaultRole, u)

			returnData := APIResponse(200, getRoles, tokenString)

			return returnData
		}
		responseData := Response(401, "Username already exists")
		return responseData
	}
	responseData := Response(401, "Email already exists")
	return responseData
}

//GetUser retrieves user data using ID
func GetUser(uid string) interface{} {
	var u User
	getUser := Conn.Where("id = ?", uid).First(&u)
	if getUser.Error != nil {
		response := Response(200, "User not exist.")

		return response
	}

	return u
}

//GetAllUsers retrieves list of all users in the system
func GetAllUsers() interface{} {
	userArray := []*User{}
	// var u User
	Conn.Find(&userArray)
	// Conn.Take(&u)
	return userArray
}

//UpdateUser updates user data
func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.FullName != "" {
			u.FullName = uu.FullName
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.DOB != "" {
			u.DOB = uu.DOB
		}
		if uu.Street != "" {
			u.Street = uu.Street
		}
		if uu.City != "" {
			u.City = uu.City
		}
		if uu.Country != "" {
			u.Country = uu.Country
		}
		if uu.Gender != "" {
			u.Gender = uu.Gender
		}
		if uu.Email != "" {
			u.Email = uu.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

//Login handles login
func Login(username, password string) (code int, user User) {
	var u User
	findEmail := Conn.Where("email = ?", username).Find(&u)
	//If Email doesn't exisit
	if findEmail != nil && findEmail.Error != nil {
		findUsername := Conn.Where("username = ?", username).Find(&u)
		if findUsername != nil && findUsername.Error != nil {
			return 404, u
		}

		if u.Password != password {
			return 401, u
		}

		return 200, u

	}
	if password != u.Password {
		return 401, u
	}

	return 200, u
}

//GetUsername gets user from username
func GetUsername(username string) User {
	var u User
	Conn.Where("username = ?", username).First(&u)
	return u
}

//GetUserEmail get user from email
func GetUserEmail(email string) User {
	var u User
	Conn.Where("email = ?", email).First(&u)
	return u
}

//DeleteUser deletes user
func DeleteUser(uid string) {
	delete(UserList, uid)
}
