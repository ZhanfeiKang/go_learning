package main

import "fmt"

/*
模拟银行账户
*/
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 1.存款
func (account *Account) Deposite(money float64, pwd string) {
	// 看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功~~")
}

// 2.取款
func (account *Account) WithDraw(money float64, pwd string) {
	// 看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	if money > account.Balance {
		fmt.Println("账户余额不足")
		return
	}

	account.Balance -= money
	fmt.Println("存款成功~~")
}

// 3.查询余额
func (account *Account) Query(pwd string) {
	// 看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	// 看看存款金额是否正确
	fmt.Println("你的账号为 :", account.AccountNo)
	fmt.Println("账户余额 :", account.Balance)
}

func main() {
	account := Account{
		AccountNo: "gs111111",
		Pwd:       "666666",
		Balance:   100.0,
	}

	account.Query("666666")
	account.Deposite(200, "666666")
	account.Query("666666")

	account.WithDraw(45.0, "666666")
	account.Query("666666")
}
