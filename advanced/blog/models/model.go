package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Category struct {
	Id   int
	Name string
}

type Posts struct {
	Id      int
	Title   string
	Info    string
	Content string
	Author  string
	Tags    string
	Created string
}

var db orm.Ormer

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8", 30)
	orm.RegisterModel(new(Category))
	db = orm.NewOrm()
}

func GetCate(Cate *[]Category) {
	o, error := orm.NewQueryBuilder("mysql")
	if error != nil {
		log.Fatal(error)
	}
	o.Select("id,name").From("categories")
	sql := o.String()
	db.Raw(sql).QueryRows(Cate)
}

func GetArticles(Articles *[]Posts) {
	o, error := orm.NewQueryBuilder("mysql")
	if error != nil {
		log.Fatal(error)
	}
	o.Select("title,info,author,tags,created").From("posts")
	sql := o.String()
	db.Raw(sql).QueryRows(Articles)
}
