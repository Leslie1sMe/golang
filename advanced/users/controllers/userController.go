package controllers

import (
	"users/model"
	"fmt"
)

type UserController struct {
	user     [] model.User
	countNum int
}

func NewUserController() *UserController {
	userController := &UserController{}
	userController.countNum = 1
	user := model.NewUser(1, "leslie", "male", 26,
		"18519268986", "239234505@qq.com")
	userController.user = append(userController.user, user)
	return userController
}

func (this *UserController) List() []model.User {
	return this.user
}

func (this *UserController) Add(user model.User) bool {
	this.countNum++
	user.Id = this.countNum
	this.user = append(this.user, user)
	return true
}

func (this *UserController) GetId(id int) int {
	num := 0
	for i := 0; i < len(this.user); i++ {
		if id == this.user[i].Id {
			num = i
		} else {
			fmt.Println("编号不合法！！！")
		}
	}
	return num
}

func (this *UserController) Del(id int) bool {
	num := this.GetId(id)
	if num == -1 {
		return false
	} else {
		this.user = append(this.user[:(num - 1)], this.user[(num + 1):]...)
		return true
	}

}

func (this *UserController) Update(id int,user model.User) bool {
	num:=this.GetId(id)
	if num == -1 {
		return false
	} else {
		this.user[num].Name = user.Name
		this.user[num].Age = user.Age
		this.user[num].Gender = user.Gender
		this.user[num].PhoneNum = user.PhoneNum
		this.user[num].Email = user.Email
		return true
	}
}
