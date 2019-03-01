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
}
