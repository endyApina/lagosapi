package controllers

import (
	"encoding/json"
	"fmt"
	"lagosapi/models"

	"github.com/astaxie/beego"
)

//OwnerController handles all operations about App Owner
type OwnerController struct {
	beego.Controller
}

//Post checks if app owner exists
// @Title AppOwner
// @Description checks if app owner exists on the system
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router /exist [get]
func (u *OwnerController) Post() {
	iAppOwner := models.DoesAppOwnerExist()
	if iAppOwner != true {
		responseData := models.UserExist(200, false)

		u.Data["json"] = responseData
		u.ServeJSON()
	}

	responseData := models.UserExist(200, true)

	u.Data["json"] = responseData
	u.ServeJSON()
}

//CreateAppOwner creates superadmin
// @Title Create App Owner
// @Description Create single app owner account.
// @Param	body		body 	models.User	true		"body for admin content"
// @Success 200 {object} models.APIResponseData
// @Failure 403 body is empty
// @router /register [post]
func (u *OwnerController) CreateAppOwner() {
	var admin models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &admin)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	if admin.Role != 999 {
		responseData := models.Response(403, "Forbidden, Not an App Owner")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	isValid := models.ValidateRegistration(admin)
	if isValid != true {
		responseData := models.Response(400, "Kindly fill all required fields")

		u.Data["json"] = responseData
		u.ServeJSON()
	}

	responseData := models.AddAppOwner(admin)
	u.Data["json"] = responseData
	u.ServeJSON()
}

//OwnerLogin function handles login for everyone.
// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {object} models.APIResponseData
// @Failure 403 user not exist
// @router /login [POST]
func (u *OwnerController) OwnerLogin() {
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
	fmt.Println(user)
	if code != 999 {
		if code == 404 {
			responseData := models.Response(404, "User Not Found")

			u.Data["json"] = responseData
			u.ServeJSON()
		}

		if code == 401 {
			responseData := models.Response(401, "Incorrect Login Details")

			u.Data["json"] = responseData
			u.ServeJSON()
		}
	}

	isAdmin := models.CheckAppOwner(user)
	if isAdmin != true {
		responseData := models.Response(403, "Unauthorized Access Point")

		u.Data["json"] = responseData
		u.ServeJSON()
	}

	getDefaultRole := models.GetRoleFromID(user.ID)
	getRoles := models.AssociateRoleUser(getDefaultRole, user)
	tokenString := models.GetToken(user)
	response := models.APIResponse(code, getRoles, tokenString)
	u.Data["json"] = response

	u.ServeJSON()
}
