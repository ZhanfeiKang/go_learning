package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": "ok",
		// })
		c.Redirect(http.StatusMovedPermanently, "http://kkite.gitee.io/")
	})

	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b 对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的URI修改
		r.HandleContext(c)        // 继续后续的处理
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "b",
		})
	})
	r.Run()
}
