package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//ValidateToken validates token
var ValidateToken = func(ctx *context.Context) {
	filter := Filter(ctx)
	if filter == true {
		return
	}

	type unAuthorized struct {
		Code int    `json:"code"`
		Body string `json:"body"`
	}

	token := ctx.Input.Header("authorization")
	validToken := ValidToken(token)
	if validToken != true {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Invalid Token"
		ctx.Output.JSON(res, false, false)

		return
	}
	if token == "" {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Invalid Token"
		ctx.Output.JSON(res, false, false)

		return
	}
	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/validate") {
		return
	}
}

//ValidToken checks if a token is valid
func ValidToken(wholeToken string) bool {
	splitString := strings.Split(wholeToken, ",")
	if splitString[0] != beego.AppConfig.String("tokenprefix") {
		return false
	}

	return true
}

//Filter checks if there are endpoint that shouldn't contain token string
func Filter(ctx *context.Context) bool {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/login") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/admin/login") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/register") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/admin/sup/register") {
		return true
	}

	return false
}
