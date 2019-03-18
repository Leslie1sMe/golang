package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Users struct {
	Id       string `orm:"unique;size(10);pk;column(id)"form:"id"`
	Username string `orm:"size(10);column(username)"form:"username"`
	Gender   string `orm:"unique;size(10);default(null);column(gender)"form:"gender"`
	Age      int    `orm:"unique;size(10);default(null);column(age)"form:"age"`
	Address  string `orm:"unique;size(20);default(null);column(address)"form:"address"`
	Email    string `orm:"unique;size(50);default(null);column(email)"form:"email"`
	Phone    string `orm:"unique;size(50);default(null);column(phone)"form:"phone"`
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

func GetAllUsers() ([]*Users, error) {
	var users []*Users
	db := orm.NewOrm()
	if _, err := db.QueryTable("users").All(&users); err != nil {
		return nil, err
	} else {
		return users, nil
	}

}

func GetOneUser(uid string) (*Users, error) {
	var user = Users{Id: uid}
	db := orm.NewOrm()
	if err := db.Read(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}

}

func UpdateUser(uid string, users Users) (*Users, error) {
	db := orm.NewOrm()
	var user = Users{Id: uid}
	if err := db.Read(&user); err != nil {
		return nil, err
	} else {
		if users.Username != "" {
			user.Username = users.Username
		}
		if users.Age != 0 {
			user.Age = users.Age
		}
		if users.Address != "" {
			user.Address = users.Address
		}
		if users.Phone != "" {
			user.Phone = users.Phone
		}
		if users.Gender != "" {
			user.Gender = users.Gender
		}
		if users.Email != "" {
			user.Email = users.Email
		}
		if _, err := db.Update(&user); err != nil {
			return nil, err
		} else {
			return &user, nil

		}
	}
}
func DeleteUser(uid string) error {
	db := orm.NewOrm()
	var user = Users{Id: uid}
	if _, err := db.Delete(&user); err != nil {
		return err
	} else {
		return nil
	}

}
