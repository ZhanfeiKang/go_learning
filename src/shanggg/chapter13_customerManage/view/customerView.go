package main

import (
	"fmt"
	"shanggg/chapter13_customerManage/model"
	"shanggg/chapter13_customerManage/service"
)

type customerView struct {
	// 定义必要字段
	key  string //接收用户输入...
	loop bool   // 表示是否循环的显示主菜单
	// 增加一个字段customerService
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (cv *customerView) list() {
	// 首先，获取到当前所有的客户信息(在切片中)
	customers := cv.customerService.List()

	// 显示
	fmt.Println("----------------------客 户 列 表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for i := 0; i < len(customers); i++ {
		// fmt.Println(customers[i].Id,"\t",customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------------------客户列表完成----------------------")
}

// add
func (cv *customerView) add() {
	fmt.Println("----------------------添 加 客 户----------------------")
	fmt.Print("姓名: ")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话: ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱: ")
	email := ""
	fmt.Scanln(&email)

	// 构建一个新的customer
	// 注意：id号没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	// 调用
	if cv.customerService.Add(customer) {
		fmt.Println("添加成功..")
	} else {
		fmt.Println("添加失败..")
	}
}

// 得到用户的输入id，删除该id对应的客户
func (cv *customerView) delete() {
	fmt.Println("----------------------删 除 客 户----------------------")
	fmt.Print("请输入待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Print("请确认是否删除(Y/N): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		// 调用 cv.customerService.Delete方法
		if cv.customerService.Delete(id) {
			fmt.Println("删除完成..")
		} else {
			fmt.Println("删除失败，输入的id号不存在..")
		}
	}
}

func (cv *customerView) update() {
	fmt.Println("----------------------修 改 客 户----------------------")
	fmt.Print("请输入待修改客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := cv.customerService.FindById(id)
	if index == -1 {
		fmt.Println("修改失败，输入的id号不存在..")
		return
	}
	customer := cv.customerService.Search(index)
	fmt.Printf("姓名(%v): ", customer.Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别(%v): ", customer.Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v): ", customer.Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话(%v): ", customer.Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v): ", customer.Email)
	email := ""
	fmt.Scanln(&email)

	// 调用 cv.customerService.Update方法
	if cv.customerService.Update(index, name, gender, age, phone, email) {
		fmt.Println("修改完成..")
	} else {
		fmt.Println("修改失败..")
	}

}

func (cv *customerView) exit() {

	fmt.Printf("确定退出用户管理系统吗(Y/N): ")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "Y" || cv.key == "n" || cv.key == "N" {
			break
		}

		fmt.Println("你的输入有误，确认是否退出(Y/N): ")
	}

	if cv.key == "y" || cv.key == "Y" {
		cv.loop = false
	}
}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println()
		fmt.Println("--------------------客户信息管理软件--------------------")
		fmt.Println("                      1 添加客户")
		fmt.Println("                      2 修改客户")
		fmt.Println("                      3 删除客户")
		fmt.Println("                      4 客户列表")
		fmt.Println("                      5 退    出")
		fmt.Print("请输入1-5: ")
		fmt.Scanln(&cv.key)
		fmt.Println()
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

		if !cv.loop {
			break
		}
		fmt.Println()
	}

	fmt.Println("你退出了客户关系管理系统..")
}

func main() {
	// 在主函数中创建一个customerView实例
	cv := customerView{
		key:  "",
		loop: true,
	}

	// 这里完成对customerView结构体的customerService字段的初始化
	cv.customerService = service.NewCustomerService()

	// 显示主菜单
	cv.mainMenu()
}
