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
// @Success 200 {object} models.ResponsePackage
// @Failure 403 body is empty
// @router /sup/ [post]
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
		responseData := models.Response(403, "Forbidden")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}
	addAdminMessage := models.AddAdmin(admin)
	responseData := models.ValidResponse(200, addAdminMessage)
	u.Data["json"] = responseData
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

// //UpdateSubAdmin updates sub admin
// // @Title UpdateSubAdmin
// // @Description update the Sub Admin
// // @Param	uid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.ResponsePackage	true		"body for user content"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is not int
// // @router /sub/:uid [put]
// func (u *AdminController) UpdateSubAdmin() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		var user models.User
// 		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 		uu, err := models.UpdateUser(uid, &user)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = uu
// 		}
// 	}
// 	u.ServeJSON()
// }

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
