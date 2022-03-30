package routers

import (
	"bubbleList/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 告诉 gin 框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateAtodoHandler)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoListHandler)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodoHandler)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodoHandler)
	}

	return r
}
