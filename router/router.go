package router

import (
	"goblog-end/config"
	"goblog-end/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.Use(config.Cors()) //开启中间件 允许使用跨域请求

	r.GET("/page/category/", controller.Category)

	r.GET("/page/userinfo/:name", controller.PageUserInfo)

	port := "8082"
	r.Run(":" + port)
}
