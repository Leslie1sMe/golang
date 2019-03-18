package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAllUsers",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOneUser",
            Router: `/one`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
