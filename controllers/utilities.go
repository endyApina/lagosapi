package controllers

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
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
	isNull := NullToken(token)
	if isNull == true {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Empty Token String"
		ctx.Output.JSON(res, false, false)

		return
	}
	if token == "" {
		var res unAuthorized
		res.Code = 403
		res.Body = "Unauthorized Connection. Empty Token"
		ctx.Output.JSON(res, false, false)

		return
	}
	isTokenExpired := TokenExpire(token)
	if isTokenExpired != true {
		var res unAuthorized
		res.Code = 401
		res.Body = "Token Expired, Kindly Login again."
		ctx.Output.JSON(res, false, false)
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

//NullToken checks if token is null
func NullToken(wholeToken string) bool {
	splitString := strings.Split(wholeToken, ",")
	if splitString[1] == "" {
		return true
	}

	return false
}

//TokenExpire checks if the user token is valid and hasn't expired
func TokenExpire(tokenS string) bool {
	wholeString := strings.Split(tokenS, ",")
	tokenString := wholeString[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwtkey")), nil
	})
	if err != nil {
		return false
	}
	var expireAt float64
	nowTime := time.Now().Add(time.Minute * 1).Unix()
	for key, val := range claims {
		if key == "expire" {
			expireAt = val.(float64)
		}
	}
	tm := float64(nowTime)
	diff := tm - expireAt
	if diff >= 360000 {
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

	if strings.HasPrefix(ctx.Input.URL(), "/v1/admin/register") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/owner/exists") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/owner/login") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/owner/register") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/admin/super/exist") {
		return true
	}

	if strings.HasPrefix(ctx.Input.URL(), "/v1/invite/validate") {
		return true
	}

	return false
}
