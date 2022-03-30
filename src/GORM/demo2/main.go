package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` // 唯一索引
	Role         string  `gorm:"size:255"`                       // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"`                // 设置会员号 (member number) 唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`                     // 给address 字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              // 忽略本字段
}

type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// 唯一指定表名
func (Animal) TableName() string {
	return "kkite"
}

func main() {
	// 修改默认的表名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:abc123@(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.SingularTable(true) // 禁用复数

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用User结构体创建名叫 xiaowangzi 的表
	// db.Table("xiaowangzi").CreateTable(&User{})
}
