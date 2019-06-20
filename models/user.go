package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	//UserList
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	// u := User{"Endy Apinageri", "astaxie", "pappi", "male", "03-05-1997", "Primewares", "Lekki", "Nigeria", "astaxie@gmail.com"}
	// UserList["user_11111"] = &u
}

//User struct shows models for users
type User struct {
	gorm.Model
	FullName string `gorm:"type:varchar(100)" json:"FullName"`
	Username string `gorm:"type:varchar(100) unique" json:"Username"`
	Password string `gorm:"type:varchar(100)" json:"Password"`
	Gender   string `gorm:"type:varchar(100)" json:"Gender"`
	DOB      string `gorm:"type:varchar(100)" json:"DOB"`
	Street   string `gorm:"type:varchar(100)" json:"Street"`
	City     string `gorm:"type:varchar(100)" json:"City"`
	Country  string `gorm:"type:varchar(100)" json:"Country"`
	Email    string `gorm:"type:varchar(100); unique_index" json:"Email"`
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
			Conn.Create(&u)

			return u
		} else {
			responseData := Response(200, "Username already exists")
			return responseData
		}
	} else {
		responseData := Response(200, "Email already exists")
		return responseData
	}
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
func Login(username, password string) interface{} {
	var u User
	findEmail := Conn.Where("email = ?", username).Find(&u)

	if findEmail != nil && findEmail.Error != nil {
		findUsername := Conn.Where("username = ?", username).Find(&u)
		if findUsername != nil && findUsername.Error != nil {
			responseData := Response(200, "Email or Username doesn't exist")
			return responseData
		}

		if u.Password != password {
			responseData := Response(200, "Incorrect Details")
			return responseData
		}

		return u

	} else {
		if u.Password != password {
			responseData := Response(200, "Incorrect Details")
			return responseData
		}

		return u
	}
}

//DeleteUser deletes user
func DeleteUser(uid string) {
	delete(UserList, uid)
}
