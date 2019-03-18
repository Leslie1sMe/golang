package controllers

import (
	"fmt"
	"github.com/Leslie1sMe/golang-sms/sms"
	"github.com/astaxie/beego"
	"github.com/monnand/goredis"
	"math/rand"
	"time"
)

type SmsController struct {
	beego.Controller
}

func (c *SmsController) Send() {
	phone := c.GetString("phone")
	fmt.Println(phone)
	var sms = sms.SendParam{"key-xxxxxxxxxxxxxxxxxxxxx", false}
	code := GetCode()
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	client.Db = 1
	if err := client.Set(phone, []byte(code)); err != nil {
		fmt.Println(err)
	}
	var message = "验证码" + code + "，您正在使用手机账号登入服务，[验证码告知他人将导致账号被盗，请勿泄漏]"
	if _, err := sms.Send(phone, message); err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = "success"
	}
	c.ServeJSON()

}

func GetCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
