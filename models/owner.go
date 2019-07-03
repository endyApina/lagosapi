package models

//AddAppOwner function adds a new super admin to the system
func AddAppOwner(a User) interface{} {
	Conn.AutoMigrate(&User{})
	Conn.AutoMigrate(&Roles{})

	appownerexist := DoesAppOwnerExist()
	if appownerexist == true {
		responseData := Response(401, "Unauthorized, App owner already exist")
		return responseData
	}

	res, user := CheckUser(a)
	if res == true {
		// Conn.Where("user_id = ?", user.ID).Delete(&Roles{})
		role := CreateDefaultRole(a)
		role.UserID = user.ID
		Conn.Create(&role)
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
	tokenString := GetTokenString(a.Username)
	getRoles := AssociateRoleUser(role, a)
	responseData := APIResponse(200, getRoles, tokenString)
	return responseData
}

//DoesAppOwnerExist checks if the app owner exists on the system
func DoesAppOwnerExist() bool {
	var r Roles
	appOwner := Conn.Where("code = 999").Find(&r)
	//if role doesn't exist
	if appOwner != nil && appOwner.Error != nil {
		return false
	}

	return true
}

//CheckAppOwner checks if a particular user is an admin
func CheckAppOwner(u User) bool {
	var r Roles
	ifAppOwner := Conn.Where("user_id = ? AND code = 999", u.ID).Find(&r)
	//if user not a sub Admin?
	if ifAppOwner != nil && ifAppOwner.Error != nil {
		return false
	}
	return true
}
