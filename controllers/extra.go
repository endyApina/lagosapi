package controllers

import (
	"lagosapi/models"

	"github.com/astaxie/beego"
)

//ExtraController handles every other system related sub-service
type ExtraController struct {
	beego.Controller
}

//ValidateToken validates token string in a request
// @Title Validate Token
// @Description validates token string in a request and returns the user information.
// @Success 200 {string} TokenString
// @router /token/validate [get]
func (u *ExtraController) ValidateToken() {
	token := u.GetString("token")
	models.ValidateToken(token)
	// u.Data["json"] = user
	// u.ServeJSON()
}
