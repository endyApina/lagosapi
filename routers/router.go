// @APIVersion 1.0.0
// @Title API's for Lagos State SME Hub
// @Description This is developed using Beego, and should be consumed by only developers in GPI
// @Contact endy.apina@my-gpi.io
// @TermsOfServiceUrl http://my-gpi.io/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"lagosapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),
		beego.NSNamespace("/token",
			beego.NSInclude(
				&controllers.TokenController{},
			),
		),
		beego.NSNamespace("/owner",
			beego.NSInclude(
				&controllers.OwnerController{},
			),
		),
		beego.NSNamespace("/invite",
			beego.NSInclude(
				&controllers.InvitationController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
