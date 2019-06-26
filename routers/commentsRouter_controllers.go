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
            Method: "GetAllIndustry",
            Router: `/`,
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
            Method: "GetAllSubAdmins",
            Router: `/sub/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:AdminController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:AdminController"],
        beego.ControllerComments{
            Method: "CreateSupAdmin",
            Router: `/sup/`,
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

    beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddIdea",
            Router: `/addidea/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lagosapi/controllers:UserController"] = append(beego.GlobalControllerRouter["lagosapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
