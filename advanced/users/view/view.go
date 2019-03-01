package main

import (
	"fmt"
	"users/controllers"
	"users/model"
)

type userView struct {
	key            string
	loop           bool
	userController *controllers.UserController
}

func main() {
	userView := userView{
		key:  "",
		loop: true,
	}
	userView.userController = controllers.NewUserController()
	//显示主菜单
	userView.showMenu()
}
func (this *userView) showMenu() {
	for {
		fmt.Println("----------用户信息管理软件-----------")
		fmt.Println("------------ 1.增 ----------------")
		fmt.Println("------------ 2.删 ----------------")
		fmt.Println("------------ 3.改 ----------------")
		fmt.Println("------------ 4.查 ----------------")
		fmt.Println("------------ 5.退 出 -------------")
		fmt.Println("请根据需要选择(1-5)")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.del()
		case "3":
			this.update()
		case "4":
			this.show()
		case "5":
			choice := ""
			for {
				fmt.Println("你确定要退出吗？y/N")
				fmt.Scanln(&choice)
				if choice == "y" || choice == "N" {
					break
				} else {
					fmt.Println("输入有误，请重新输入！！！")
				}
			}
			if choice == "y" {
				this.loop = false
				fmt.Println("感谢使用")
			}

		default:
			fmt.Println("输入有误请重新输入！！！")
		}
		if !this.loop {
			break
		}

	}
}

func (this *userView) add() {
	fmt.Println("----------------用户列表----------------")
	fmt.Println("请输入姓名")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("请输入性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("请输入年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("请输入手机号")
	phoneNum := ""
	fmt.Scanln(&phoneNum)
	fmt.Println("请输入邮箱")
	email := ""
	fmt.Scanln(&email)
	user := model.NewUser2(name, gender, age, phoneNum, email)
	this.userController.Add(user)
	fmt.Println("创建成功！！")

}

func (this *userView) show() {

	user := this.userController.List()
	if len(user) == 0 {
		fmt.Println("当前无数据，快添加一条吧！！！")
	}
	fmt.Println("----------------用户列表----------------")
	fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t手机号\t\t邮箱\t\t")
	for i := 0; i < len(user); i++ {
		fmt.Println(user[i].ShowUser())
	}
	fmt.Println()
	this.userController.List()
}

func (this *userView) del() {
	fmt.Println("请输入要删除的数据编号(输入-1退出)")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	if this.userController.Del(id) {
		fmt.Println("删除成功！！！")
	} else {
		fmt.Println("删除失败！！！")
	}
}

func (this *userView) update() {
	fmt.Println("请输入要修改的数据编号(输入-1退出)")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	num:=this.userController.GetId(id)
		fmt.Println("请重新输入名字")
		name:=""
		fmt.Scanln(&name)
		fmt.Println("请重新输入性别")
		gender:=""
		fmt.Scanln(&gender)
		fmt.Println("请重新输入年龄")
		age:=0
		fmt.Scanln(&age)
		fmt.Println("请重新输入手机号")
		phoneNum:=""
		fmt.Scanln(&phoneNum)
		fmt.Println("请重新输入邮箱")
		email:=""
		fmt.Scanln(&email)
		user := model.NewUser2(name,gender,age,phoneNum,email)
		if this.userController.Update(num,user){
			fmt.Println("修改成功")
		}else {
			fmt.Println("修改失败")
		}

}
