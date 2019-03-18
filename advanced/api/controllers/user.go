package controllers

import (
	"api/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/monnand/goredis"
)

type UserController struct {
	beego.Controller
}

// @Title Add A User
// @Description create user
// @Param username query string false "username"
// @Param gender query string false "gender"
// @Param age query int false "age"
// @Param address query string false "address"
// @Param email query string false "email"
// @Param phone query string false "phone"
// @Success 200 {string} uid
// @Failure  400 {error} err
// @router /add [post]
func (u *UserController) AddUser() {
	var user models.Users
	if err := u.ParseForm(&user); err != nil {
		u.Data["json"] = err
	} else {
		uid := models.AddUser(&user)
		u.Data["json"] = uid
	}

	u.ServeJSON()
}

// @Title Get All Users
// @Description Get All The Users
// @Success 200 {object} []*Users
// @Failure  400 {error} err
// @router /all [get]
func (u *UserController) GetAllUsers() {
	if res, err := models.GetAllUsers(); err != nil {
		u.Data["json"] = err
		fmt.Println(err)
	} else {
		u.Data["json"] = res
	}
	u.ServeJSON()
}

// @Title Get One Users
// @Description Get  The Users
// @Success 200 {object} *Users
// @Param uid query string false "id"
// @Failure  400 {error} err
// @router /one [get]
func (u *UserController) GetOneUser() {
	var uid string = u.GetString("uid")
	if res, err := models.GetOneUser(uid); err != nil {
		u.Data["json"] = err
	} else {
		fmt.Println(res)
		u.Data["json"] = res
	}
	u.ServeJSON()
}

// @Title Update A User
// @Description Update user
// @Param uid query string false "id"
// @Param username query string false "username"
// @Param gender query string false "gender"
// @Param age query int false "age"
// @Param address query string false "address"
// @Param email query string false "email"
// @Success 200 {object} *Users
// @Failure  400 {error} err
// @router /update [post]
func (u *UserController) UpdateUser() {
	var user models.Users
	u.ParseForm(&user)
	if res, err := models.UpdateUser(user.Id, user); err != nil {
		u.Data["json"] = err
	} else {
		fmt.Println(res)
		u.Data["json"] = res
	}
	u.ServeJSON()
}

// @Title Delete One Users
// @Description Get  The Users
// @Success 200 {string} "delete success"
// @Param uid query string false "id"
// @Failure  400 {error} err
// @router /delete [get]
func (u *UserController) DeleteUser() {
	var uid = u.GetString("uid")
	if err := models.DeleteUser(uid); err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = " delete success"
	}
	u.ServeJSON()
}

// @Title Delete One Users
// @Description Get  The Users
// @Success 200 {string} "{\"status_code\":\"200\",\"message\":\"login success\"}"
// @Param phone query string false "phone"
// @Param code query string false
// @Failure  400 {string} "{\"status_code\":\"400\",\"message\":\"login failed\"}"
// @router /login [post]
func (u *UserController) Login() {
	code := u.Input().Get("code")
	phone := u.Input().Get("phone")
	var client goredis.Client
	client.Db = 1
	if redisCode, err := client.Get(phone); err != nil {
		u.Data["json"] = err
	} else {
		if code != string(redisCode) {
			u.Data["json"] = "{\"status_code\":\"400\",\"message\":\"login failed\"}"
		} else {
			u.Data["json"] = "{\"status_code\":\"200\",\"message\":\"login success\"}"
		}
	}
	u.ServeJSON()
}
