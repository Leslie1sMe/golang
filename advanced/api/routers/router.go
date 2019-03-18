// @APIVersion 1.0.0
// @Title beego API
// @Author Leslie
// @Description beego useful api demo
// @Contact lizhonglin_melody@yeah.net
// @Paths api/controllers
package routers

import (
	"api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/add", &controllers.UserController{}, "post:AddUser"),
			beego.NSRouter("/all", &controllers.UserController{}, "get:GetAllUsers"),
			beego.NSRouter("/one", &controllers.UserController{}, "get:GetOneUser"),
			beego.NSRouter("/update", &controllers.UserController{}, "post:UpdateUser"),
			beego.NSRouter("/delete", &controllers.UserController{}, "get:DeleteUser"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/sms",
			beego.NSRouter("/send", &controllers.SmsController{}, "get:Send"),
			beego.NSInclude(
				&controllers.SmsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
