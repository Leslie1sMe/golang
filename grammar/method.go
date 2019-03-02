package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

type Circle struct {
	radius float64
}

func (this Person) speak() {
	fmt.Println(this.Name + "是个大好人")
}

func (this Person) count() {
	y := 0
	for x := 1; x <= 1000; x++ {
		y += x
	}
	println(y)
}

func (this Person) count1(n int) {
	y := 0
	for x := 1; x <= n; x++ {
		y += x
	}
	println(y)
}

func (this Person) getSum(x int, y int) {
	fmt.Println(x + y)
}

func (this Circle) area() float64 {
	return this.radius * this.radius * math.Pi
}

type UserController struct {
}

type MethodUtils struct {
}

type Method struct {
	x     int
	y     int
	shape string
}

type Method1 struct {
	x     float64
	y     float64
	shape string
}

func main() {
	//var person Person
	//person.Name = "张扬"
	//person.speak()
	//person.count()
	//var n int
	//fmt.Scanln(&n)
	//person.count1(n)
	//var a int
	//var b int
	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	//person.getSum(a,b)

	//var circle  = Circle{3.0}
	//res:=circle.area()
	//fmt.Println(res)
	var user MethodUtils
	user.Print()
	user.Print2(10, 2)
	area := user.area(10.4, 2.3)
	fmt.Println(area)
	user.judge(9)
	var test Method
	test.test(10, 9, "@")
	var x Method1
	res := x.test1(1.0, 2.4, "@")
	fmt.Println(res)
}

func (user MethodUtils) Print() {
	for x := 1; x <= 10; x++ {
		for y := 1; y <= 8; y++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func (user MethodUtils) Print2(m int, n int) {
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func (user MethodUtils) area(width float64, length float64) float64 {
	return width * length
}
func (user MethodUtils) judge(num int) {
	if num%2 == 0 {
		fmt.Printf("%v是偶数", num)
	} else {
		fmt.Printf("%v是奇数", num)

	}
}

func (method Method) test(m int, n int, shape string) {
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			fmt.Print(shape)
		}
		fmt.Println("")
	}
}

func (method Method1) test1(m float64, n float64, operator string) (float64) {
	res:=0.0
	switch operator {
	case "+":
		res  = m + n
		return res
	case "-":
		res = m - n
		return res

	case "*":
		res = m * n
		return res

	case "/":
		res = m / n
		return res

	default:
		fmt.Println("操作符有误")
		return res
	}
}
