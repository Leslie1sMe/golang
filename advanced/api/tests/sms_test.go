/*
* @Author:Leslie
* @Date: 2019/3/18 16:11
 */
package test

import (
	"github.com/Leslie1sMe/golang-sms/sms"
	"testing"
)

func TestSms(t *testing.T) {
	var s = sms.SendParam{"key-xxxxxxxxxxxxxxxxxxxx", false}
	if res, err := s.Send("185xxxxx8986", "验证码222222，您正在使用手机账号登入服务，[验证码告知他人将导致账号被盗，请勿泄漏]"); err != nil {
		t.Log("wtf")

	} else {
		t.Log(res)
	}
}
