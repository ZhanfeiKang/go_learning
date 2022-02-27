package model

import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account {

	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对..")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码的长度不对..")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额需足够20元..")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

func (account *account) SetPwd(pwd string) {
	account.pwd = pwd
	// 判断密码长度是否符合要求
	// 等等打印
}

// ... ...
