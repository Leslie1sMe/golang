package main

import "fmt"

type Usb interface {
	Starting()
	Shutdown()
}
type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

func (computer Computer) work(usb Usb) {
	usb.Starting()
	if phone, ok := usb.(Phone); ok {
		phone.Calling()
	}
	usb.Shutdown()
}
func (c Camera) Starting() {
	fmt.Println("照相机设备启动")
}

func (c Camera) Shutdown() {
	fmt.Println("照相机设备关闭")
}
func (phone Phone) Calling() {
	fmt.Println("手机设备打电话")
}
func (phone Phone) Starting() {
	fmt.Println("手机设备启动")
}

func (phone Phone) Shutdown() {
	fmt.Println("手机设备关闭")
}

func main() {
	//var a interface{}
	//var j int64
	//a = j
	//b := a.(int32)
	//fmt.Println(b)
	//断言类型失败，报错panic，为避免panic错误，我们对程序进行改写
	//var a interface{}
	//var j int32
	//a = j
	//b, ok := a.(int64)
	//if !ok {
	//	fmt.Println("类型断言失败")
	//} else {
	//	fmt.Println("类型断言成功")
	//	fmt.Println(b)
	//
	//}
	//现在我们换一种数据结构试一下
	//var a interface{}
	//b := B{"Leslie", 22}
	//a = b
	//c, ok := a.(B)
	//if !ok {
	//	fmt.Println("断言失败")
	//} else {
	//	fmt.Println("断言成功")
	//	fmt.Println(c)
	//}

	//var a int32
	//var b int64
	//c := "Hello"
	//var h float64
	//var i float32
	//d := B{"Leslie", 20}
	//f := &B{"Leslie", 20}
	//TypeJudge(a, b, c, d, f, h, i)

	var b [3]Usb
	var computer Computer
	b[0] = Phone{"xiaomi"}
	b[1] = Camera{"canna"}
	b[2] = Phone{"apple"}
	for _, v := range b {
		computer.work(v)
	}

}

//func TypeJudge(a ...interface{}) {
//	for num, b := range a {
//		switch b.(type) {
//		case float64:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		case float32:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		case int64:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		case int32:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		case string:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		case B:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		case *B:
//			fmt.Printf("第%d个参数的值为%v\n", num+1, b)
//		default:
//			fmt.Println("第%d个参数的类型未知", num)
//		}
//	}
//}
