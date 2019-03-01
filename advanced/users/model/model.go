package model

import "fmt"

type User struct {
	Id       int
	Name     string
	Gender   string
	Age      int
	PhoneNum string
	Email    string
}

func NewUser(id int, name string, gender string,
	age int, phoneNum string, email string) User {
	return User{
		Id:       id,
		Name:     name,
		Gender:   gender,
		Age:      age,
		PhoneNum: phoneNum,
		Email:    email,
	}
}
func NewUser2(name string, gender string,
	age int, phoneNum string, email string) User {
	return User{
		Name:     name,
		Gender:   gender,
		Age:      age,
		PhoneNum: phoneNum,
		Email:    email,
	}
}

func (this User) ShowUser() string {
	data := fmt.Sprintf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t", this.Id, this.Name,
		this.Gender, this.Age, this.PhoneNum, this.Email)
	return data
}
