package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 灵活使用tag来对结构体字段做定制化操作
type msg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法1：使用map
		// data := map[string]interface{}{
		// 	"name": "小王子",
		// 	"msg":  "hello world~",
		// 	"age":  18,
		// }

		// 方法2
		data := gin.H{
			"name": "小王子",
			"msg":  "hello world~",
			"age":  18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "小王子",
			Message: "Hello golang",
			Age:     18,
		}
		c.JSON(http.StatusOK, data) // json的序列化
	})

	r.Run(":8080")
}
