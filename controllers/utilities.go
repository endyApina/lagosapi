package controllers

import (
	"log"

	"github.com/astaxie/beego/context"
)

//ValidateToken validates token
var ValidateToken = func(ctx *context.Context) {
	// token := ctx.Request.FormValue("token")
	// token := ctx.Input.Header("token")
	// ip := ctx.Request.Header.Get("X-Forwarded-For")
	ip := ctx.Request.RemoteAddr
	log.Println(ip)
}

func CheckIPAddress(ip string) bool {
	if ip != "127.0.0.1" {
		return false
	}

	return true
}
