package controllers

import (
	"encoding/json"
	"lagosapi/models"

	"github.com/astaxie/beego"
)

//UserController handles all operations about Users
type UserController struct {
	beego.Controller
}

//Post create New Users
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.APIResponseData
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Post() {
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	isValid := models.ValidateRegistration(user)
	if isValid != true {
		responseData := models.Response(400, "Kindly fill all required fields")

		u.Data["json"] = responseData
		u.ServeJSON()
	}
	addUserMessage := models.AddUser(user)
	u.Data["json"] = addUserMessage
	u.ServeJSON()
}

//GetAll gets all Users in the System
// @Title GetAll
// @Description get all Users
// @Success 200 {object} []models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	responseData := models.ValidResponse(200, users)
	u.Data["json"] = responseData
	u.ServeJSON()
}

//Get retrieves user data by ID
// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ResponsePackage
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	getUserMessage := models.GetUser(uid)
	responseData := models.ValidResponse(200, getUserMessage)
	u.Data["json"] = responseData
	u.ServeJSON()
}

// //Put Update user data
// // @Title Update
// // @Description update the user
// // @Param	uid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {object} models.ResponsePackage
// // @Failure 403 :uid is not int
// // @router /:uid [put]
// func (u *UserController) Put() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		var user models.User
// 		err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 		if err != nil {
// 			responseData := models.Response(200, "Invalid JSON format")

// 			u.Data["json"] = responseData
// 			u.ServeJSON()

// 			return
// 		}
// 		uu, err := models.UpdateUser(uid, &user)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = uu
// 		}
// 	}
// 	u.ServeJSON()
// }

// //Delete removes user from the system
// // @Title Delete
// // @Description delete the user
// // @Param	uid		path 	string	true		"The uid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 uid is empty
// // @router /:uid [delete]
// func (u *UserController) Delete() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }

//Login function handles login for everyone.
// @Title Login
// @Description Logs user into the system
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.APIResponseData
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var loginDetails models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &loginDetails)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	username := loginDetails.Username
	password := loginDetails.Password

	code, user := models.Login(username, password)
	if code != 200 {
		if code == 404 {
			responseData := models.Response(404, "User does not exist.")

			u.Data["json"] = responseData
			u.ServeJSON()
		}

		if code == 401 {
			responseData := models.Response(401, "Invalid login credentials.")

			u.Data["json"] = responseData
			u.ServeJSON()
		}
	}
	isAdmin := models.CheckAdmin(user)
	if isAdmin == true {
		responseData := models.Response(403, "Unauthorized Access Point")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}

	getDefaultRole := models.CreateDefaultRole(user)
	getRoles := models.AssociateRoleUser(getDefaultRole, user)
	tokenString := models.GetTokenString(username)
	response := models.APIResponse(code, getRoles, tokenString)
	u.Data["json"] = response
	u.ServeJSON()
}
