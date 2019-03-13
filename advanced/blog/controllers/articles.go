package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) List() {
	var Articles []models.Posts
	models.GetArticles(&Articles)
	c.Data["posts"] = Articles
	c.TplName = "list.tpl"
	//var Context *context.Context
	//paginator := pagination.SetPaginator(Context, 10, int64(len(Articles)))
	//c.Data["gt"] = paginator
}

func (c *ArticleController) Detail() {
	s := c.Ctx.Input.Param(":id")
	Id, error := strconv.Atoi(s)
	if error != nil {
		log.Fatal(error)
	}
	c.EnableRender = true
	c.TplName = "detail.tpl"
	Post, err := models.GetArticle(Id)
	if err != nil {
		log.Fatal(err)
	}
	c.Data["created"] = Post.Created
	c.Data["id"] = Post.Id
	c.Data["tags"] = Post.Tags
	c.Data["title"] = Post.Title
	c.Data["author"] = Post.Author
	c.Data["content"] = Post.Content
	c.Data["info"] = Post.Info
	if Id > 1 {
		PostFront, _ := models.GetArticle(Id - 1)
		c.Data["front"] = PostFront
		c.Data["is_first"] = false
	} else {
		c.Data["is_first"] = true
	}
	PostBack, _ := models.GetArticle(Id + 1)
	c.Data["backend"] = PostBack
}
