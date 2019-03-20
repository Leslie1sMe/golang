/*
* @Author:Leslie
* @Date: 2019/3/20 14:10
 */
package main

import (
	"go_code/spy"
)

func main() {
	spy.GetToWork("http://tieba.baidu.com/f?kw=科比&ie=utf-8&pn=50", `data-original="(?s:(.*?))"`)
}
