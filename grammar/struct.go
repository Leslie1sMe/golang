//结构体的操作练习
// struct和其他语言的class具有同等地位
//golang没有extends继承是通过匿名字段
//面向对象编程=>面向接口编程
package main

type User struct {
	Name   string
	Age    int
	Gender string
	Score  [3]int
}

type Cat struct {
	Name     string
	Color    string
	Age      int
	Behavior string
}
type Struct struct {
	Per    *int
	Name   string
	Age    int
	Score  [] float64
	Hobby  [3] float64
	Friend map[string]string
}


func main() {
	//user := User{
	//	Name:"Leslie",
	//	Age:27,
	//	Gender:"male",
	//}
	//println(user.Gender)
	//修改结构体字段的value
	//user := User{
	//	Name:"Leslie",
	//	Age:27,
	//	Gender:"male",
	//}
	//fmt.Println(user.Gender)
	//user.Age = 30
	//fmt.Println(user)
	//cat1:=Cat{
	//	Name:"小白",
	//	Age:10,
	//	Color:"白色",
	//}
	//cat2:=Cat{
	//	Name:"小花",
	//	Age:20,
	//	Color:"花色",
	//}
	//fmt.Println(cat1)
	//fmt.Println(cat2)
	//struct中包含array
	//var user User
	//user.Score = [3]int{100,100,100}
	//user.Name = "Leslie"
	//fmt.Println(user)
	//var p1 Struct
	//fmt.Println(p1)
	//{<nil>  0 [] [0 0 0] map[]}
	//if p1.Per == nil {
	//	fmt.Println("true")
	//}
	//struct中包含slice
	//p1.Score = append(p1.Score,149.9)
	//p1.Score = make([]float64,10)
	//p1.Score[0] = 10
	//fmt.Println(p1)
	//struct中包含map
	//p1.Friend = make(map[string]string)
	//p1.Friend["Name"] = "NC"
	//fmt.Println(p1)
	//结构体为值类型，默认为值拷贝
	//创造结构体变量的四种方式
	//var U User
	//var U User = User{}
	//U:=User{}
	//var U *User = new(User)
	//	var person *User = &User{"Leslie",19,"male",[3]int{1,2,3}}
	//(*U).Name  = "Leslie"
	//(*U).Score = [...]int{100,20,30}
	//fmt.Println(*U)
	////简化：
	//U.Gender = "male"
	//fmt.Println(U)
	//var person *User = &User{"Leslie",19,"male",[3]int{1,2,3}}
	//fmt.Println(person)


}



//test()和UserController绑定
//test()只能由UserController类型的变量调用
