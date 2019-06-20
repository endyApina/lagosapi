package controllers

import (
	"encoding/json"
	"lagosapi/models"
	"log"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

// @Title CreateAdmin
// @Description Create more administrative users.
// @Param	body		body 	models.Admin	true		"body for admin content"
// @Success 200 {int} models.Admin.Id
// @Failure 403 body is empty
// @router / [post]
func (u *AdminController) Post() {
	var admin models.Admin
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &admin)
	if err != nil {
		log.Println(err)
	}
	addAdminMessage := models.AddAdmin(admin)
	responseData := models.ValidResponse(200, addAdminMessage)
	u.Data["json"] = responseData
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *AdminController) GetAll() {
	users := models.GetAllAdmin()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *AdminController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetAdmin(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *AdminController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *AdminController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
