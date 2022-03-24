package model

import (
	"fmt"
	"testing"
)

// TestMain 函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	// 通过 m.Run() 来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	// 通过 t.Run() 来执行子测试函数
	// t.Run("测试添加用户: ", testAddUser)
	// t.Run("测试获取用户: ", testGetUserById)
	t.Run("测试获取所有用户: ", testGetUsers)
}

// 如果函数名不是以Test开头，那么该函数默认不执行，我们可以将它设置成为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("子测试函数执行")
	// fmt.Println("测试添加用户：")
	// user := &User{}
	// user.AddUser()
	// user.AddUser2()
}

// 测试获取一个User
func testGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录: ")
	user := &User{
		ID: 1,
	}

	// 调用获取User的方法
	u, _ := user.GetUerById()
	fmt.Println("得到的User的信息是：", u)
}

// 测试获取所有User
func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录: ")
	user := &User{}
	// 调用获取所有User的方法
	us, _ := user.GetUsers()
	// 遍历切片
	for k, u := range us {
		fmt.Printf("第%d个用户：%v\n", k+1, u)
	}
}
