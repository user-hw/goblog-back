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

	r.GET("/post/", controller.Post)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Avatar":   "/resource/images/3.jpg",
			"UserName": "不会写代码的建筑师",
			"UserDesc": "写代码是医生的邀请",
			"Bilibili": "https://space.bilibili.com/66628849",
			"zhihu":    "https://www.zhihu.com/people/xu-heng-wei-63",
			"github":   "https://github.com/user-hw",
			"code":     200,
			// "Categorys": [1,2],
		})

	})

	port := "8082"
	r.Run(":" + port)
}
