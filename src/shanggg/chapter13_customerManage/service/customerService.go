package service

import (
	"shanggg/chapter13_customerManage/model"
)

// 该 CustomerService ,完成对Customer的操作，包括
// 增删改查
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户
	// 该字段后面，还可以作为新客户的id+1
	customerNum int
}

// 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "11223244322", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// 添加客户到customers切片
// ！！！一定要是指针的绑定
func (cs *CustomerService) Add(customer model.Customer) bool {
	// 我们确定一个id分配的规则，就是添加的顺序
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// 根据id删除客户
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	// index == -1,说明没有这个客户
	if index == -1 {
		return false
	}

	// 如何从切片中删除一个元素
	// [:index] : [0,index)
	// cs.customers[index+1:]...        ... 表示展开切片传入
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// 通过id查找客户，并返回一个customer
func (cs *CustomerService) Search(index int) model.Customer {
	return cs.customers[index]
}

// 修改客户信息
func (cs *CustomerService) Update(index int, name string,
	gender string, age int, phone string, email string) bool {

	// 修改客户信息
	if name != "" {
		cs.customers[index].Name = name
	}
	if gender != "" {
		cs.customers[index].Gender = gender
	}
	if age != 0 {
		cs.customers[index].Age = age
	}
	if phone != "" {
		cs.customers[index].Phone = phone
	}
	if email != "" {
		cs.customers[index].Email = email
	}

	return true

}

// 根据id查找客户在切片中对应下标，如果没有该客户，返回-1
func (cs *CustomerService) FindById(id int) int {
	index := -1
	// 遍历cs.customers 切片
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			// 找到
			index = i
			return index
		}
	}
	return index
}
