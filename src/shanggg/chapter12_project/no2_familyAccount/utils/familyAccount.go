package utils

import "fmt"

type FamilyAccount struct {
	// 声明一个变量，保存接收用户输入的选项
	key string

	// 声明一个变量，控制是否退出for循环
	loop bool

	// 定义账户余额
	blance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 定义一个变量，记录是否有收支的行为
	flag bool
	// 收支的详情使用字符串来记录
	// 当有收支时，只需要对details进行拼接处理即可
	details string
}

// 编写工厂模式的构造方法，返回一个 *FamilyAccount 实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		blance:  10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}

// 将显示明细写成一个方法
func (family *FamilyAccount) showDetails() {
	fmt.Println("-----------------------当前收支明细记录-----------------------")
	if family.flag {
		fmt.Println(family.details)
	} else {
		fmt.Println("目前没有任何收支，请来一笔吧~")
	}
}

// 登记收入
func (family *FamilyAccount) income() {
	fmt.Print("本次收入金额: ")
	fmt.Scanln(&family.money)
	family.blance += family.money // 修改账户余额
	fmt.Print("本次收入说明: ")
	fmt.Scanln(&family.note)
	// 将这个收入情况拼接到details变量
	family.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", family.blance, family.money, family.note)
	family.flag = true
}

func (family *FamilyAccount) pay() {
	fmt.Print("本次支出金额: ")
	fmt.Scanln(&family.money)
	// 这里需要做一个必要的判断
	if family.money > family.blance {
		fmt.Println("余额的金额不足")
		return
	}
	family.blance -= family.money
	fmt.Print("本次支出说明: ")
	fmt.Scanln(&family.note)
	family.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", family.blance, family.money, family.note)
	family.flag = true
}

func (family *FamilyAccount) exit() {
	fmt.Print("你确定要退出吗？(y/n): ")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("你你的输入有误，请重新输入(y/n): ")
	}
	if choice == "y" {
		family.loop = false
	}
}

// 给该结构体绑定响应的方法
// 显示主菜单
func (family *FamilyAccount) MainMenu() {
	// 显示这个主菜单
	for {
		fmt.Println("\n-----------------------家庭收支记账软件-----------------------")
		fmt.Println("                         1 收支明细")
		fmt.Println("                         2 登记收入")
		fmt.Println("                         3 登记支出")
		fmt.Println("                         4 退出软件")

		fmt.Print("请选择(1-4): ")

		fmt.Scanln(&family.key)
		switch family.key {
		case "1":
			family.showDetails()
		case "2":
			family.income()
		case "3":
			family.pay()
		case "4":
			family.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !family.loop {
			break
		}
	}
	fmt.Println("你已退出了家庭记账软件的使用")
}
