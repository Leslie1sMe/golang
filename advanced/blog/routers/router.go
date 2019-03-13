package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	beego.Router("/list", &controllers.ArticleController{}, "get:List")
	beego.Router("/detail/:id:int", &controllers.ArticleController{}, "get:Detail")
}
