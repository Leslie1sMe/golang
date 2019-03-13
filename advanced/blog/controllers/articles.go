package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) List() {
	var Articles []models.Posts
	models.GetArticles(&Articles)
	c.Data["posts"] = Articles
	c.TplName = "list.tpl"
}

func (c *ArticleController) Detail() {
	c.TplName = "detail.tpl"
}
