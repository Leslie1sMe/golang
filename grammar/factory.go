package main

import (
	"fmt"
	"go_code/grammar/model"
)

//使用工厂模式操作model
func main() {
	fmt.Println("golang使用工厂模式！")
	student := model.NewStudent("Leslie", 120)
	fmt.Println(student)

}
