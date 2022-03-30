package main

// 警告：
// 	删除记录时，请确保主键字段有值，GORM会通过主键去删除记录，如果主键为空，GORM会删除该model的所有记录。

// 如果不喜欢然删除，在定义结构体的时候不加入Model就可以了

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
	// u1 := User{Name: "kkite2", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu2", Age: 20, Active: false}
	// db.Create(&u2)

	// 5.删除
	// var u User
	// // u.ID = 1
	// u.Name = "jinzhu2" // 会全部删除，开头警告
	// db.Debug().Delete(&u)

	// db.Debug().Where("name=?", "jinzhu2").Delete(User{})
	// db.Delete(User{}, "age=?", 18)

	// var u1 []User
	// db.Debug().Unscoped().Where("name=?", "jinzhu2").Find(&u1)
	// fmt.Println(u1)

	// 物理删除
	db.Debug().Unscoped().Where("name=?", "kkite2").Delete(User{})
}
