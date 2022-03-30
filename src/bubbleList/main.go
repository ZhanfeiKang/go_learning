package main

import (
	"bubbleList/dao"
	"bubbleList/models"
	"bubbleList/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;

	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出去关闭数据库连接

	// 绑定模型
	dao.DB.AutoMigrate(&models.Todo{}) // todos

	r := routers.SetupRouter()

	r.Run(":8080")
}
