package main

import "fmt"

type Account struct {
	AccountNo string
	Password  string
	Balance   float64
}

func (account *Account) Deposite(money float64, password string) {

	if password != account.Password {
		fmt.Println("密码错误！！！")
	}
	if money <= 0 {
		fmt.Println("输入金额有误！！！")
	}
	account.Balance += money
	fmt.Println("存款成功！！！")

}

func (account *Account) WithDraw(money float64, password string) {
	if password != account.Password {
		fmt.Println("密码错误！！！")
	}
	if money > account.Balance {
		fmt.Println("输入金额有误！！！")
	}
	account.Balance -= money
	fmt.Println("取款成功！！！")

}

func (account Account) Query(password string) {
	if password != account.Password {
		fmt.Println("密码错误！！！")
	}
	fmt.Println(account.Balance)
}

func main() {
	//account := Account{
	//	AccountNo: "xxxxxxxxxx",
	//	Password:  "123456",
	//	Balance:   100000.0,
	//}

	//var account Account
	//account.Balance = 10000.0
	//password := ""
	//fmt.Scanln(&password)
	//account.Password = password
	//account.Query(account.Password)

}
