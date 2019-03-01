package account

import "fmt"

type PrivateAccount struct {
	choice  int
	loop    bool
	balance float64
	income  float64
	expense float64
	info    string
	detail  string
	flag    bool
}

func NewPrivateAccount() *PrivateAccount {
	return &PrivateAccount{
		choice:  1,
		loop:    false,
		income:  0.0,
		expense: 0.0,
		info:    "",
		flag:    false,
		balance: 0.0,
		detail:  "收支 / 账号余额 / 收支金额 / 说明 /t",
	}
}

func (this *PrivateAccount) MainMenu() {
	for {
		fmt.Println("----------个人收支记录软件-----------")
		fmt.Println("------------ 1.收入详情 -------------")
		fmt.Println("------------ 2.登记收入 -------------")
		fmt.Println("------------ 3.登记支出 -------------")
		fmt.Println("------------ 4.退出程序 -------------")
		fmt.Println("请根据需要选择(1-4)")
		fmt.Scanln(&this.choice)
		switch this.choice {
		case 1:
			this.AccountInfo()
		case 2:
			this.IncomeAccount()
		case 3:
			this.ExpenseAccount()
		case 4:
			this.Exit()
		default:
			fmt.Println("操作无效，请重新输入！")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("感谢使用！！！")

}

func (this *PrivateAccount) AccountInfo() {
	if !this.flag {
		fmt.Println("当前没有任何收支情况哦!")
	} else {
		fmt.Println("----------- 收支明细 ------------")
		fmt.Println(this.detail)
	}
	this.loop = true
}

func (this *PrivateAccount) IncomeAccount() {
	fmt.Println("----------- 本次收入金额 ------------")
	fmt.Scanln(&this.income)
	this.balance += this.income
	fmt.Println("----------- 本次收入说明 ------------")
	fmt.Scanln(&this.info)
	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.income, this.info)
	this.flag = true
	this.loop = true

}
func (this *PrivateAccount) ExpenseAccount() {
	fmt.Println("----------- 本次支出金额  ------------")
	fmt.Scanln(&this.expense)
	if this.balance < this.expense {
		fmt.Println("余额不足,不能支出！")
	} else {
		this.balance -= this.expense
		fmt.Println("----------- 本次支出说明  ------------")
		fmt.Scanln(&this.info)
		this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.expense, this.info)
		this.flag = true
	}
	this.loop = true

}

func (this *PrivateAccount) Exit() {
	decision := ""
	fmt.Println("您确认要退出吗？y/N")
	for {
		fmt.Scanln(&decision)
		if decision == "y" || decision == "N" {
			break
		}else{
			fmt.Println("输入无效请重新输入y/N")
		}
	}
	if decision == "y" {
		this.loop = false
	}
	if decision == "N" {
		this.loop = true
	}


}
