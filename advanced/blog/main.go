package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"time"
)

func GetTime() string {
	s := time.Now().Format("2006-01-02 15:04:05")
	return s
}

func main() {
	beego.AddFuncMap("d", GetTime)
	beego.SetStaticPath("/detail", ".")
	beego.Run()
}
