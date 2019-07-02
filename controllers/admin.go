package controllers

import (
	"encoding/json"
	"lagosapi/models"

	"github.com/astaxie/beego"
)

//AdminController handles admin controller
type AdminController struct {
	beego.Controller
}

//CreateSupAdmin creates superadmin
// @Title CreateAdmin
// @Description Create more super administrative users.
// @Param	body		body 	models.User	true		"body for admin content"
// @Success 200 {object} models.APIResponseData
// @Failure 403 body is empty
// @router /register [post]
func (u *AdminController) CreateSupAdmin() {
	var admin models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &admin)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	if admin.Role != 99 {
		responseData := models.Response(403, "Forbidden, User not an Admin")

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

	responseData := models.AddAdmin(admin)
	u.Data["json"] = responseData
	u.ServeJSON()
}

//AdminLogin function handles login for everyone.
// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {object} models.APIResponseData
// @Failure 403 user not exist
// @router /login [POST]
func (u *AdminController) AdminLogin() {
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
			responseData := models.Response(404, "User Not Found")

			u.Data["json"] = responseData
			u.ServeJSON()
		}

		if code == 401 {
			responseData := models.Response(401, "Incorrect Details")

			u.Data["json"] = responseData
			u.ServeJSON()
		}
	}

	isAdmin := models.CheckAdmin(user)
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

//InviteSubAdmin function invites sub admins to the system by sending an email link.
// @Title Invite Sub Admin
// @Description Invites other admin users to join the system
// @Param	body		body 	models.Invite	true		"A json containing the role {int}, email {string}, tokenString"
// @Success 200 {string} invitation sent!
// @Failure 403 user not exist
// @router /invite [post]
func (u *AdminController) InviteSubAdmin() {
	var invite models.Invite
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &invite)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}

	sendInvitation := models.SpecialInvite(invite)
	u.Data["json"] = sendInvitation
	u.ServeJSON()
}

//CreateSubAdmin creates sub admin
// @Title SubAdmin
// @Description Create more sub Admins .
// @Param	body		body 	models.APIData	true		"body for admin content"
// @Success 200 {object} models.ResponsePackage
// @Failure 403 body is empty
// @router /sub/ [post]
func (u *AdminController) CreateSubAdmin() {
	var body models.APIData
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &body)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	createMessage := models.CreateSubAdmin(body)
	responseData := models.ValidResponse(200, createMessage)
	u.Data["json"] = responseData
	u.ServeJSON()
}

//GetAll gets all
// @Title GetAll
// @Description get all Users
// @Success 200 {object} []models.ResponseObject
// @router / [get]
func (u *AdminController) GetAll() {
	users := models.GetAllAdmin()
	u.Data["json"] = users
	u.ServeJSON()
}

//GetSupAdmin gets all superior admin
// @Title GetSupAdmin
// @Description Get all Super Administrative users
// @Success 200 {object} models.User
// @router /sup/ [get]
func (u *AdminController) GetSupAdmin() {
	users := models.GetAllSupAdmin()
	if users == nil {
		responseData := models.Response(204, "No Super Admin")
		u.Data["json"] = responseData
		u.ServeJSON()
	}
	responseData := models.ValidResponse(200, users)
	u.Data["json"] = responseData
	u.ServeJSON()
}

//GetAllSubAdmins gets all subadmins
// @Title AllSubAdmins
// @Description get all Sub Admins
// @Success 200 {object} []models.ResponseObject
// @router /sub/ [get]
func (u *AdminController) GetAllSubAdmins() {
	users := models.GetAllAdmin()
	u.Data["json"] = users
	u.ServeJSON()
}

//SuperAdminExist checks if a super admin exist or not
// @Title SuperAdmin
// @Description checks if super admin exists on the system
// @Param	body		body 	models.User	"body for user content"
// @Success 200 {object} models.APIResponseData
// @Failure 403 body is empty
// @router /super/exist [get]
func (u *AdminController) SuperAdminExist() {
	iAppOwner := models.DoesSupAdminExist()
	if iAppOwner != true {
		responseData := models.UserExist(200, false)

		u.Data["json"] = responseData
		u.ServeJSON()
	}

	responseData := models.UserExist(200, true)

	u.Data["json"] = responseData
	u.ServeJSON()
}

// //DeleteSubAdmin deletes sub admin
// // @Title DeleteSubAdmin
// // @Description delete the sub admin
// // @Param	uid		path 	string	true		"The uid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 uid is empty
// // @router /sub/:uid [delete]
// func (u *AdminController) DeleteSubAdmin() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }
