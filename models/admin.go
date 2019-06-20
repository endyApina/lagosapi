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

func AddAdmin(a Admin) interface{} {
	Conn.AutoMigrate(&Admin{})

	findEmail := Conn.Where("username = ?", a.Username).Find(&a)
	if findEmail != nil && findEmail.Error != nil {
		addAdmin := Conn.Create(&a)
		if addAdmin != nil && addAdmin.Error != nil {
			panic(addAdmin.Error)
		}

		return a

	}
	return 0
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
