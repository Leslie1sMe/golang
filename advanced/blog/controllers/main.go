package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var Cate []models.Category
	models.GetCate(&Cate)
	var Articles []models.Posts
	models.GetArticles(&Articles)
	c.Data["articles"] = Articles
	c.Data["m"] = Cate
	c.TplName = "index.tpl"

}

func (c *MainController) Contact() {
	c.TplName = "contact.tpl"
	c.Data["author"] = "Leslie"
	c.Data["professor"] = "Backend Engineer"
}

func (c *MainController) About() {
	c.TplName = "about.tpl"
}
