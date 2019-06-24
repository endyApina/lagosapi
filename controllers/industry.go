package controllers

import (
	"lagosapi/models"

	"github.com/astaxie/beego"
)

//IndustryController handles admin controller
type IndustryController struct {
	beego.Controller
}

//GetAll gets all Industry in the database
// @Title GetAll
// @Description get all Industry in the database
// @Success 200 {object} []models.ResponseObject
// @router / [get]
func (u *AdminController) GetAllIndustry() {
	users := models.GetAllAdmin()
	u.Data["json"] = users
	u.ServeJSON()
}
