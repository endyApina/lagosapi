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
            Method: "CreateSupAdmin",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "GetAllSubAdmins",
            Router: `/sub/`,
            AllowHTTPMethods: []string{"get"},
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
            Method: "SuperAdminExist",
            Router: `/super/exist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:OwnerController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:OwnerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/exist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:OwnerController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:OwnerController"],
        beego.ControllerComments{
            Method: "CreateAppOwner",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:TokenController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:TokenController"],
        beego.ControllerComments{
            Method: "ValidateToken",
            Router: `/validate`,
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

}
