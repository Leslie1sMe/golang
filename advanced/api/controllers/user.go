package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) AddUser() {
	var user models.Users
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		u.Data["json"] = err
	} else {
		fmt.Println(user)
	}
	//uid := models.AddUser(&user)
	//data, _ := json.Marshal(uid)
	//u.Data["json"] = data
	//u.Data["status"] = 200
	u.ServeJSON()
}

//
//func (u *UserController) Post() {
//	var user models.User
//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//	uid := models.AddUser(user)
//	u.Data["json"] = map[string]string{"uid": uid}
//	u.ServeJSON()
//}
//
//func (u *UserController) GetAll() {
//	users := models.GetAllUsers()
//	u.Data["json"] = users
//	u.ServeJSON()
//}
//
//func (u *UserController) Get() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		user, err := models.GetUser(uid)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = user
//		}
//	}
//	u.ServeJSON()
//}
//
//func (u *UserController) Put() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(uid, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}
//
//func (u *UserController) Delete() {
//	uid := u.GetString(":uid")
//	models.DeleteUser(uid)
//	u.Data["json"] = "delete success!"
//	u.ServeJSON()
//}
//
//func (u *UserController) Login() {
//	username := u.GetString("username")
//	code := u.GetString("password")
//	if models.Login(username, code) {
//		u.Data["json"] = "dadsa"
//	} else {
//		u.Data["json"] = "adasdasdada"
//	}
//	u.ServeJSON()
//}
//
//func (u *UserController) Logout() {
//	u.Data["json"] = "logout success"
//	u.ServeJSON()
//}
