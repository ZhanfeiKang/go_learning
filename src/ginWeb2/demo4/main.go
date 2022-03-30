package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取form表单提交的参数

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// /login post
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单数据
		// username := c.PostForm("username")
		// password := c.PostForm("password")

		// username := c.DefaultPostForm("username", "somebody")
		// password := c.DefaultPostForm("xxx", "***")

		username, ok := c.GetPostForm("username")
		if !ok {
			username = "kk"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":8080")
}
