package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var (
	UserList map[string]*Users
)

type Users struct {
	Id       string `orm:"unique;size(10);pk;column(id)"json:"id"`
	Username string `orm:"size(10);column(username)"json:"username"`
	Password string `orm:"unique;size(20);column(password)"json:"password"`
	Gender   string `orm:"unique;size(10);default(null);column(gender)"json:"gender"`
	Age      int    `orm:"unique;size(10);default(null);column(age)"json:"age"`
	Address  string `orm:"unique;size(20);default(null);column(address)"json:"address"`
	Email    string `orm:"unique;size(20);default(null);column(email)"json:"email"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func AddUser(u *Users) string {
	db := orm.NewOrm()
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	_, err := db.Insert(u)
	if err != nil {
		panic(err)
	}
	return u.Id
}

//
//func GetUser(uid string) (u *User, err error) {
//	if u, ok := UserList[uid]; ok {
//		return u, nil
//	}
//	return nil, errors.New("User not exists")
//}
//
//func GetAllUsers() map[string]*User {
//	return UserList
//}
//
//func UpdateUser(uid string, uu *User) (a *User, err error) {
//	if u, ok := UserList[uid]; ok {
//		if uu.Username != "" {
//			u.Username = uu.Username
//		}
//		if uu.Password != "" {
//			u.Password = uu.Password
//		}
//		if uu.Profile.Age != 0 {
//			u.Profile.Age = uu.Profile.Age
//		}
//		if uu.Profile.Address != "" {
//			u.Profile.Address = uu.Profile.Address
//		}
//		if uu.Profile.Gender != "" {
//			u.Profile.Gender = uu.Profile.Gender
//		}
//		if uu.Profile.Email != "" {
//			u.Profile.Email = uu.Profile.Email
//		}
//		return u, nil
//	}
//	return nil, errors.New("User Not Exist")
//}
//
//func Login(username, password string) bool {
//	for _, u := range UserList {
//		if u.Username == username && u.Password == password {
//			return true
//		}
//	}
//	return false
//}
//
//func DeleteUser(uid string) {
//	delete(UserList, uid)
//}
