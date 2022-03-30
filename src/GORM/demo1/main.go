package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接数据库
	db, err := gorm.Open("mysql", "root:abc123@(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移 （把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	// u1 := UserInfo{
	// 	ID:     1,
	// 	Name:   "kkite",
	// 	Gender: "男",
	// 	Hobby:  "篮球",
	// }
	// db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u) // 查询表中第一条数据，保存到u中
	fmt.Printf("u:%#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "see daisy")

	// 删除
	db.Delete(&u)
}
