package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "InviteSubAdmin",
            Router: `/invite`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "AdminLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAllSubAdmins",
            Router: `/sub/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "CreateSubAdmin",
            Router: `/sub/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetSupAdmin",
            Router: `/sup/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "CreateSupAdmin",
            Router: `/sup/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:ExtraController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:ExtraController"],
        beego.ControllerComments{
            Method: "ValidateToken",
            Router: `/token/validate`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "ValidateToken",
            Router: `/validate`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
