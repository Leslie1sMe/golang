package main

import "fmt"

//golang的interface test
type Animals interface {
	walk()
	eat()
}

type animals struct {
}

type Tiger struct {
}

type Dog struct {
}

func (d Dog) walk() {
	fmt.Println(" is walking")
}

func (d Dog) eat() {
	fmt.Println(" is eating")
}

func (t Tiger) walk() {
	fmt.Println(" is walking")
}

func (t Tiger) eat() {
	fmt.Println(" is eating")
}

func (a animals) Dosometing(animals Animals) {
	animals.eat()
	animals.walk()
}

func main() {
	d := Dog{}
	c := Tiger{}
	a := animals{}
	a.Dosometing(d)
	a.Dosometing(c)

}
