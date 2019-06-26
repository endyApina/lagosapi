package controllers

import (
	"log"

	"github.com/astaxie/beego/context"
)

//ValidateToken validates token
var ValidateToken = func(ctx *context.Context) {
	log.Println("Work")
}
