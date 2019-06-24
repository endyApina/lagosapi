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
// @Success 200 {object} models.ResponsePackage
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		responseData := models.Response(200, "Invalid JSON format")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	addUserMessage := models.AddUser(user)
	responseData := models.ValidResponse(200, addUserMessage)
	u.Data["json"] = responseData
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

//Put Update user data
// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.ResponsePackage
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		if err != nil {
			responseData := models.Response(200, "Invalid JSON format")

			u.Data["json"] = responseData
			u.ServeJSON()

			return
		}
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

//Delete removes user from the system
// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

//Login function handles login for everyone.
// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")

	loginMessage := models.Login(username, password)
	responseData := models.ValidResponse(200, loginMessage)
	u.Data["json"] = responseData

	u.ServeJSON()
}

//Logout handles logout
// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

//AddIdea adds a new business idea to the user profile
// @Title AddIdea
// @Description create new Ideas
// @Param	body		body 	models.APIData	true		"body to add new Business Idea"
// @Success 200 {object} models.ResponsePackage
// @Failure 403 body is empty
// @router /addidea [post]
func (u *UserController) AddIdea() {
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		responseData := models.Response(200, "Invalid JSON format")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	addUserMessage := models.AddUser(user)
	responseData := models.ValidResponse(200, addUserMessage)
	u.Data["json"] = responseData
	u.ServeJSON()
}
