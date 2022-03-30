package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1.定义模型
type User struct {
	gorm.Model // ID CreatedAt  UpdatedAt  DeletedAt
	Name       string
	Age        int64
	Active     bool
}

func main() {
	// 2.连接数据库
	db, err := gorm.Open("mysql", "root:abc123@(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 3.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 4.创建
	// u1 := User{Name: "kkite", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu", Age: 20, Active: false}
	// db.Create(&u2)

	// 5.查询
	var user User
	db.First(&user)

	// 6.更新
	user.Name = "daisy"
	user.Age = 3
	// db.Debug().Save(&user) // 会更新所有字段

	// db.Model(&user).Update("name", "Daisy") // 更新某个字段

	// m1 := map[string]interface{}{
	// 	"name":   "kkite",
	// 	"age":    22,
	// 	"active": true,
	// }
	// db.Debug().Model(&user).Updates(m1)                // 更新多个指定字段
	// db.Debug().Model(&user).Select("age").Update(m1)   // 只更新age字段
	// db.Debug().Model(&user).Omit("active").Updates(m1) // 排除m1中的active，更新其他字段

	// db.Debug().Model(&user).UpdateColumn("age", 30)
	// rowsNum := db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected // 查询了多少行
	// fmt.Println(rowsNum)

	// 让users表中所有的用户的年龄在原来的基础上+2
	db.Model(&User{}).Update("age", gorm.Expr("age+?", 2))
}
