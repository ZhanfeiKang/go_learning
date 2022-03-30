package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
type User struct {
	gorm.Model // ID CreatedAt  UpdatedAt  DeletedAt
	Name       string
	Age        int64
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:abc123@(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	// u1 := User{Name: "kkite", Age: 18}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu", Age: 20}
	// db.Create(&u2)

	// 4.查询
	// var user User // 声明模型结构体类型变量user
	// user := new(User)
	// db.First(user)
	// fmt.Printf("user:%#v\n", user)

	// fmt.Println("-----------------------------------------")

	// var users []User
	// db.Debug().Find(&users)
	// fmt.Printf("users:%#v\n", users)

	// FirstOrInit
	var user User
	// 未找到
	db.FirstOrInit(&user, User{Name: "小王子"})
	fmt.Printf("user:%#v\n", user)

	// 找到
	db.FirstOrInit(&user, User{Name: "jinzhu"})
	fmt.Printf("user:%#v\n", user)

	// 未找到就将attr的值赋给user
	db.Attrs(User{Age: 99}).FirstOrInit(&user, User{Name: "小王子"})
	fmt.Printf("user:%#v\n", user)

	// 不管找没找到都用99去更新
	db.Assign(User{Age: 99}).FirstOrInit(&user, User{Name: "jinzhu"})
	fmt.Printf("user:%#v\n", user)
}
