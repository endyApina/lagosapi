package models

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/jinzhu/gorm"
)

var (
	//AdminList is houses an array of Admin objects
	AdminList map[string]*Admin
)

func init() {
	AdminList = make(map[string]*Admin)
	// u := User{"Endy Apinageri", "astaxie", "pappi", "male", "03-05-1997", "Primewares", "Lekki", "Nigeria", "astaxie@gmail.com"}
	// UserList["user_11111"] = &u
}

//Admin struct holds admin object
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

//AddAdmin function adds a new super admin to the system
func AddAdmin(a User) interface{} {
	Conn.AutoMigrate(&User{})
	Conn.AutoMigrate(&Roles{})

	supadminexist := DoesSupAdminExist()
	if supadminexist == true {
		responseData := Response(401, "Unauthorized, Super Admin already exist")
		return responseData
	}

	res, user := CheckUser(a)
	if res == true {
		// Conn.Where("user_id = ?", user.ID).Delete(&Roles{})
		role := CreateDefaultRole(a)
		role.UserID = user.ID
		Conn.Create(&role)
		SendRegistrationEmail(a)
		tokenString := GetTokenString(a.Username)
		getRoles := AssociateRoleUser(role, a)
		responseData := APIResponse(200, getRoles, tokenString)
		return responseData
	}

	checkUsername := CheckUsername(a)
	if checkUsername == true {
		responseData := Response(403, "Username already Exists")
		return responseData
	}

	hashPassword, _ := HashPassword(a.Password)
	a.Password = hashPassword
	Conn.Create(&a)
	role := CreateDefaultRole(a)
	Conn.Create(&role)
	SendRegistrationEmail(a)
	tokenString := GetToken(a)
	getRoles := AssociateRoleUser(role, a)
	responseData := APIResponse(200, getRoles, tokenString)
	return responseData
}

//CreateSubAdmin creates a new admin to the system
func CreateSubAdmin(api APIData) interface{} {
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

//GetAdmin returns the information of admin with uid
func GetAdmin(uid string) (u *Admin, err error) {
	if u, ok := AdminList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Admin not exists")
}

//GetAllAdmin returns list of all admins
func GetAllAdmin() interface{} {
	var rolesArray []Roles
	Conn.Where("code = 99").Or("code = 88").Find(&rolesArray)

	allAdmin := GetUserFromRole(rolesArray)
	return allAdmin
}

//GetAllSupAdmin gets all super admin
func GetAllSupAdmin() []*Roles {
	var roleArray []*Roles
	Conn.Where("code = 99").Find(&roleArray)

	return roleArray
}

//UpdateAdmin edits or update admin records.
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

//Invitation stores invitation user object
type Invitation struct {
	gorm.Model
	Email            string `gorm:"type:varchar(100)" json:"email"`
	Role             int    `gorm:"type:int(100)" json:"role"`
	VerificationCode string `gorm:"type:varchar(100)" json:"code"`
}

//SpecialInvite sends an invitation link to whoever email is passed
func SpecialInvite(invite Invitation) interface{} {
	var u User
	u.Email = invite.Email
	u.Role = invite.Role
	userType := GetUserType(invite.Role)
	template := beego.AppConfig.String("templatepath") + "invite.html"

	res, user := CheckUser(u)
	verifi := StoreInvite(user, invite.Role)
	if res != true {
		mailLink := beego.AppConfig.String("currentip") + "special/register/" + u.Email + "/" + strconv.Itoa(invite.Role) + "/" + verifi
		user.FullName = strings.Split(user.Email, "@")[0]
		SendInviteMessage(user, mailLink, template, userType)

		responseData := Response(200, "Mail Sent Successfully")
		return responseData
	}

	mailLink := beego.AppConfig.String("currentip") + "special/login/" + u.Email + "/" + strconv.Itoa(invite.Role) + "/" + verifi
	user.FullName = strings.Split(user.Email, "@")[0]
	SendInviteMessage(user, mailLink, template, userType)

	responseData := Response(200, "Invitation sent successfully")
	return responseData
}

//StoreInvite function stores invitation details in the db
func StoreInvite(u User, role int) string {
	Conn.AutoMigrate(&Invitation{})
	var invitation Invitation
	invitation.Email = u.Email
	invitation.Role = role
	invitation.VerificationCode = GenerateRandCode(16)

	invitationExist := CheckInvite(invitation)
	if invitationExist == false {
		Conn.Create(&invitation)

		return invitation.VerificationCode
	}
	Conn.Model(&invitation).Update("verification_code", invitation.VerificationCode)
	return invitation.VerificationCode
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "!£$%^&*()[]<>"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//StringWithCharset chill
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//GenerateRandCode gets new random codes
func GenerateRandCode(length int) string {
	return StringWithCharset(length, charset)
}

//CheckInvite checks if theres an invite for a particular for that same role.
func CheckInvite(invite Invitation) bool {
	var i Invitation
	invitation := Conn.Where("email = ? AND role = ?", invite.Email, invite.Role).Find(&i)
	//If invitation does not exist
	if invitation != nil && invitation.Error != nil {
		return false
	}

	return true
}
