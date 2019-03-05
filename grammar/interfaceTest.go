package main

import "fmt"

//继承多个接口
//****接口是引用类型
//type A interface {
//	sayA()
//	B
//	C
//}
//
//type B interface {
//	sayB()
//}
//
//type C interface {
//	sayC()
//}
//
//type Test struct {
//}
//
//func (test Test) sayA() {
//	fmt.Println("sayA")
//}
//func (test Test) sayB() {
//	fmt.Println("sayB")
//}
//func (test Test) sayC() {
//	fmt.Println("sayC")
//}
//func main() {
//	var test Test
//	var a A = test
//	a.sayB()
//	a.sayA()
//	a.sayC()
//}
//
//type A interface {
//	Say()
//}
//
//type Test struct {
//}
//
//func (t *Test) Say() {
//	fmt.Println("wocao")
//}
//
//func main() {
//	var test Test
//	var a A = &test
//	a.Say()
//	fmt.Println(a)
//}
//
//type Test struct {
//	Name string
//	Age  int
//}
//
//type TestSlice []Test
//
//func (test TestSlice) Len() int {
//	return len(test)
//}
//
//func (test TestSlice) Less(i, j int) bool {
//	return test[i].Age < test[j].Age
//}
//
//func (test TestSlice) Swap(i, j int) {
//	//temp := test[i]
//	//test[i] = test[j]
//	//test[j] = temp
//	test[i], test[j] = test[j], test[i]
//
//}
//
//func main() {
//	//slice := []int{1, 3, 4, 2, -1, 9}
//	//sort.Ints(slice)
//	//fmt.Println(slice)
//	var slice TestSlice
//	for i := 0; i < 10; i++ {
//		test := Test{
//			Name: fmt.Sprintf("测试%d", rand.Intn(100)),
//			Age:  rand.Intn(100),
//		}
//		slice = append(slice, test)
//	}
//	fmt.Println(slice)
//	sort.Sort(slice)
//	for _, v := range slice {
//		fmt.Println(v)
//	}
//}
//
//type Student1 struct {
//	Name  string
//	Age   int
//	Score float64
//}
//
//type StudentSlice []Student1
//
//func (s StudentSlice) Len() int {
//	return len(s)
//}
//
//func (s StudentSlice) Less(i, j int) bool {
//	return s[i].Score > s[j].Score
//}
//
//func (s StudentSlice) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func main() {
//	var s StudentSlice
//	for n := 0; n < 10; n++ {
//		student := Student1{
//			Name:  fmt.Sprintf("Leslie的儿子%d", rand.Intn(19)),
//			Age:   rand.Intn(1993),
//			Score: rand.Float64(),
//		}
//		s = append(s, student)
//	}
//	fmt.Println(s)
//	sort.Sort(s)
//	for _, v := range s {
//		fmt.Println(v)
//	}
//
//}

type Monkey struct {
	Name string
}

func (this Monkey) climbing() {
	fmt.Println(this.Name + "生来会爬树")
}

type LittleMonkey struct {
	Monkey //继承
}

//不破坏原来老猴子的规则和结构，扩展出新的功能，使用实现接口的方式。
//A结构体继承了B结构体，那么A结构体自动继承了B结构体的字段和方法，并且可以直接使用
//当A结构体需要扩展功能，而又不希望破坏继承关系，则可以采用实现接口的方法
type Learn interface {
	flying()
	running()
}

func (m *LittleMonkey) flying() {
	fmt.Println(m.Name + "is flying")
}

func (m *LittleMonkey) running() {
	fmt.Println(m.Name + "is running")
}

func main() {
	var a LittleMonkey
	a.Name = "孙悟空"
	a.climbing()
	var b Learn = &a
	b.flying()
	b.running()
}
