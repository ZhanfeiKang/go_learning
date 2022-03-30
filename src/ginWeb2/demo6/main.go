package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// u := UserInfo{
		// 	username: username,
		// 	password: password,
		// }
		var u UserInfo          // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u) // ?
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo          // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u) // ?
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo          // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u) // ?
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	})

	r.Run(":8080")
}
