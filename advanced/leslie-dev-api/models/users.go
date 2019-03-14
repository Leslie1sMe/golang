package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Uid       int    `orm:"column(id) auto"`
	UserName  string `orm:"column(user_name);size(30);null"`
	PhoneNum  string `orm:"column(phone_num);size(50);null"`
	Email     string `orm:"column(email);size(30);null"`
	Avatar    string `orm:"column(avatar);size(30);http://api.pupupula.cn/uploads/applogo108108.png"`
	Badge     int    `orm:"column(badge);size(10);0"`
	DeviceId  string `orm:"column(device_id);size(50);null"`
	AreaCode  string `orm:"column(area_code);size(20);86"`
	CreatedAt string `orm:"column(created_at);size(30);null"`
	UpdatedAt string `orm:"column(updated_at);size(30);null"`
}

func (c *Users) TableName() string {
	return "categories"
}

var db orm.Ormer

func init() {
	orm.RegisterModel(new(Users))
	db = orm.NewOrm()
}

//增加
func Add(m *Users) (num int, err error) {
	o := orm.NewOrm()
	o.Insert(&m)
	return
}

//改
func Update(m *Users) {
	o := orm.NewOrm()
	c := Users{Uid: m.Uid}
	if err := o.Read(&c); err == nil {
		if n, err := o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", n)
		}
	}
}

//查询得到一条数据
func Get(Id int) (c *Users, err error) {
	o := orm.NewOrm()
	c = &Users{Uid: Id}
	if err := o.Read(c); err == nil {
		return c, nil
	}
	return nil, err

}

//查询得到数据
var user map[int]Users

func GetAll(querything string, field string, table string, order string) {
	var m Users
	o, _ := orm.NewQueryBuilder("mysql")
	o.Select(querything).Where(field + "=?").From(table).OrderBy(order)
	sql := o.String()
	db.Raw(sql).QueryRows(m)

}

//删
func Delete(Id int) {
	o := orm.NewOrm()
	var c = Users{Uid: Id}
	if err := o.Read(&c); err == nil {
		if n, err := o.Delete(&Users{Uid: Id}); err == nil {
			fmt.Println("Number of records deleted in database:", n)
		}
	}

}
