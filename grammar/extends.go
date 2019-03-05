package main

import "fmt"

type Student struct {
	Name  string
	Score float64
}

//结构体可以使用匿名结构体的任意字段（大写或者小写均可使用）或者方法
type Graduate struct {
	Student
	Gender string
}
type Pupil struct {
	Student
	Age int
}

//继承的用法
//代码复用性提高
//可维护性和可扩展性提高了

func main() {
	//p := Pupil{Student{"Leslie", 100.0}, 26}
	//g := Graduate{Student{"Leslie", 100.0}, "male"}
	//p.test()
	//p.Student.SetScore(200.3)
	//score1 := p.Student.GetScore()
	//fmt.Println(score1)
	//g.test()
	//g.Student.SetScore(149.5)
	//score2 := g.Student.GetScore()
	//fmt.Println(score2)
}

func (s *Student) GetScore() float64 {
	return s.Score
}

func (s *Student) SetScore(score float64) {
	s.Score = score
}

func (p *Pupil) test() {
	fmt.Println("Pupils are testing")
}

func (g *Graduate) test() {
	fmt.Println("Graduates are testing")
}
