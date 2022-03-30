package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由与路由组

func main() {
	r := gin.Default()

	// 访问 /index 的GET请求，会走这一条处理逻辑
	// 路由
	// r.HEAD()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	// Any: 请求方法的大集合/大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			// ...
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"method": "Any",
		// })
	})
	// NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"method": "kkite.com",
		})
	})

	// 视频的首页和详情页
	// r.GET("/video/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/index",
	// 	})
	// })
	// r.GET("/video/aa", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/aa",
	// 	})
	// })
	// r.GET("/video/bb", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/bb",
	// 	})
	// })
	// 路由组的组  多用于区分不同的业务线或API版本
	// 把公用的前缀提出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/aa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/aa",
			})
		})
		videoGroup.GET("/bb", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/bb",
			})
		})
	}

	// 商城的首页和详情页
	// r.GET("/shop/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/shop/index",
	// 	})
	// })
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/index",
			})
		})
		shopGroup.GET("/aa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/aa",
			})
		})
		shopGroup.GET("/bb", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/bb",
			})
		})
	}

	r.Run(":9090")
}
