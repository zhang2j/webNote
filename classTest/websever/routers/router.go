package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users map[string]string = make(map[string]string)

func LoadRouters(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data":   "Hello World!",
		})
	})

	//router.POST("/login", controllers.LoginTest)
	router.GET("/login", func(c *gin.Context) {
		fmt.Println("调用login方法...")
		var username string = c.Query("username")
		var password string = c.Query("password")

		fmt.Println(1, username, password)

		value, ok := users[username]

		if ok && value == password {
			fmt.Println("返回ok...")
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "ok",
			})
		} else {
			fmt.Println("返回err...")
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "err",
			})
		}
	})

	router.GET("/register", func(c *gin.Context) {
		fmt.Println("调用register方法...")
		var username string = c.Query("username")
		var password string = c.Query("password")

		fmt.Println(1, username, password)

		_, ok := users[username]

		if ok {
			fmt.Println("用户名已存在...")
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "username has been registered",
			})
		} else {
			fmt.Println("可以注册")
			users[username] = password
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "successful",
			})
		}
	})
}
