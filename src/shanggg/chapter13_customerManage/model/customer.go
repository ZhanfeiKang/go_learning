package model

import "fmt"

// 声明一个Customer结构体，表示一个客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 编写一个工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 第二种创建customer实例方法，不带id
// 编写一个工厂模式，返回一个Customer的实例
func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 返回用户的信息,格式化后的字符串
func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t\t%v", customer.Id,
		customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)

	return info
}
