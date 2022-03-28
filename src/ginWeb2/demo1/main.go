package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 静态文件
// html 页面上用到的样式文件 .css js文件 图片
func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "./statics")
	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		// HTTP 请求
		c.HTML(http.StatusOK, "posts/index.html", gin.H{ // 模板渲染
			"title": "https://kkite.gitee.io/",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		// HTTP 请求
		c.HTML(http.StatusOK, "users/index.html", gin.H{ // 模板渲染
			"title": "<a href='https://kkite.gitee.io/'>kkite's blog</a>",
		})
	})

	r.Run(":8080") // 启动server
}
