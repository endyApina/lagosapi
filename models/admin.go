package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	AdminList map[string]*Admin
)

func init() {
	AdminList = make(map[string]*Admin)
	// u := User{"Endy Apinageri", "astaxie", "pappi", "male", "03-05-1997", "Primewares", "Lekki", "Nigeria", "astaxie@gmail.com"}
	// UserList["user_11111"] = &u
}

type Admin struct {
	gorm.Model
	FullName string `gorm:"type:varchar(100)" json:"FullName"`
	Username string `gorm:"type:varchar(100)" json:"Username"`
	Password string `gorm:"type:varchar(100)" json:"Password"`
	Gender   string `gorm:"type:varchar(100)" json:"Gender"`
	DOB      string `gorm:"type:varchar(100)" json:"DOB"`
	Street   string `gorm:"type:varchar(100)" json:"Street"`
	City     string `gorm:"type:varchar(100)" json:"City"`
	Country  string `gorm:"type:varchar(100)" json:"Country"`
	Email    string `gorm:"type:varchar(100)" json:"Email"`
}

func AddAdmin(a User) interface{} {
	Conn.AutoMigrate(&User{})
	Conn.AutoMigrate(&Roles{})
	var r Roles

	findEmail := Conn.Where("email = ?", a.Email).Find(&a)

	//If email doesn't exist, proceed to check if username exist
	if findEmail != nil && findEmail.Error != nil {
		checkUsername := Conn.Where("username = ?", a.Username).Find(&a)
		if checkUsername != nil && checkUsername.Error != nil {
			//If Email and Username doesn't exist.
			if a.Role == 99 {
				checkSuperAdmin := Conn.Where("code = ?", a.Role).Find(&r)
				if checkSuperAdmin != nil && checkSuperAdmin.Error != nil {
					Conn.Create(&a)
					role := CreateDefaultRole(a)
					Conn.Create(&role)

					responseData := AssociateRoleUser(role, a)
					return responseData
				} else {
					responseData := Response(403, "Super Admin Already Exists")
					return responseData
				}
			}

			responseData := Response(403, "Forbidden")
			return responseData
		} else {
			responseData := Response(403, "Username already exists")
			return responseData
		}
	} else {
		responseData := Response(403, "Email already exists")
		return responseData
	}
}

func CreateSubAdmin(api ApiData) interface{} {
	admin := api.User
	sub := api.Body

	checkUsername := CheckUsername(sub)
	if checkUsername != false {
		responseData := Response(403, "Username already exists")
		return responseData
	}

	ifAdminExists := CheckUserExists(admin)
	if ifAdminExists != true {
		responseData := Response(403, "Admin Doesn't exist")
		return responseData
	}

	ifSuperAdmin := CheckSuperAdmin(admin)
	if ifSuperAdmin != true {
		responseData := Response(401, "Unauthorized, Cannot create Admin")
		return responseData
	}

	ifSubAdminExists := CheckUserExists(sub)
	if ifSubAdminExists == true {
		if sub.Role != 88 {
			responseData := Response(403, "Forbidden, Not a Sub Admin")
			return responseData
		}

		ifSubAdmin := CheckSubAdmin(sub)
		if ifSubAdmin == true {
			responseData := Response(403, "Forbidden, User already a Sub Admin")
			return responseData
		}

		role := CreateDefaultRole(sub)
		Conn.Create(&role)
		responseData := GetUserRoles(sub)
		return responseData
	}

	Conn.Create(&sub)
	role := CreateDefaultRole(sub)
	Conn.Create(&role)

	responseData := GetUserRoles(sub)
	return responseData

}

func GetAdmin(uid string) (u *Admin, err error) {
	if u, ok := AdminList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Admin not exists")
}

func GetAllAdmin() map[string]*Admin {
	return AdminList
}

func UpdateAdmin(uid string, uu *Admin) (a *Admin, err error) {
	if u, ok := AdminList[uid]; ok {
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
