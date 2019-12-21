package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_learn/routers"
)

func main() {
	r := gin.Default()

	// 设置跨域
	r.Use(cors.Default())

	// 加载路由
	routers.LoadRouters(r)

	r.Static("/static", "./static")

	r.Run(":8081")
}
