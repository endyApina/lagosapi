package controllers

import (
	"encoding/json"
	"lagosapi/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//TokenController handles every endpoint relating to token
type TokenController struct {
	beego.Controller
}

//InvitationController handles every the endpoint relating to inviting users
type InvitationController struct {
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
	var email string
	for key, val := range claims {
		if key == "username" {
			username = val.(string)
		}

		if key == "email" {
			email = val.(string)
		}
	}

	var user models.User
	if username != "" {
		user = models.GetUsername(username)
	}

	if email != "" {
		user = models.GetUserEmail(email)
	}

	getDefaultRole := models.CreateDefaultRole(user)
	getRoles := models.AssociateRoleUser(getDefaultRole, user)
	response := models.APIResponse(200, getRoles, tokenString)
	u.Data["json"] = response
	u.ServeJSON()
}

//VerifyInvite verifies a user invitation
// @Title Verify Invitation Link
// @Description validates a user invitation link to see if it was actually created by user.
// @Param	body		body 	models.Invitation	true		"A json containing the role {int}, email {string} and code {string}"
// @Success 200 {string} "Invitation Url"
// @router /validate [POST]
func (u *InvitationController) VerifyInvite() {
	var invite models.Invitation
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &invite)
	if err != nil {
		responseData := models.Response(405, "Method Not Allowed")

		u.Data["json"] = responseData
		u.ServeJSON()

		return
	}

	code := models.VerifyInvite(invite)
	if code != 200 {
		if code == 404 {
			responseData := models.Response(404, "User does not exist.")

			u.Data["json"] = responseData
			u.ServeJSON()
		}

		if code == 403 {
			responseData := models.Response(403, "Forbidden, Invalid Invitation Link.")

			u.Data["json"] = responseData
			u.ServeJSON()
		}
	}

	user := models.GetUserEmail(invite.Email)
	getDefaultRole := models.CreateDefaultRole(user)
	getRoles := models.AssociateRoleUser(getDefaultRole, user)
	tokenString := models.GetToken(user)
	response := models.APIResponse(code, getRoles, tokenString)
	u.Data["json"] = response
	u.ServeJSON()
}
