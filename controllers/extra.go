package controllers

import (
	"lagosapi/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//TokenController handles every other system related sub-service
type TokenController struct {
	beego.Controller
}

//ValidateToken validates token string in a request
// @Title Validate Token
// @Description validates token string in a request and returns the user information.
// @Success 200 {string} TokenString
// @router /validate [get]
func (u *TokenController) ValidateToken() {
	tokenS := u.Ctx.Input.Header("authorization")
	wholeString := strings.Split(tokenS, ",")
	if wholeString[0] != beego.AppConfig.String("tokenprefix") {
		responseData := models.Response(403, "Invalid token")

		u.Data["json"] = responseData
		u.ServeJSON()
	}
	tokenString := wholeString[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	if err != nil {
		responseData := models.Response(403, "Invalid token")

		u.Data["json"] = responseData
		u.ServeJSON()
	}
	var username string
	for key, val := range claims {
		if key == "secret" {
			username = val.(string)
		}
	}
	user := models.GetUsername(username)
	getDefaultRole := models.CreateDefaultRole(user)
	getRoles := models.AssociateRoleUser(getDefaultRole, user)
	response := models.APIResponse(200, getRoles, tokenString)
	u.Data["json"] = response
	u.ServeJSON()
}
