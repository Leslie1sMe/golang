package model

type Student struct {
	Name string
	Age  int
}

//适用于结构体首字母小写的情况，可使用工厂模式
func New(name string, age int) *Student {
	//return &Student{
	//	Name: name,
	//	Age:  age,
	//}
	return &Student{
		Name: name,
		Age:  age,
	}
}
