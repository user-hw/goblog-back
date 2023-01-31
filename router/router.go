/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2022-12-21 16:32:01
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 21:59:36
 * @FilePath: /goblog-back/router/router.go
 * @Description: 路由文件
 */
package router

import (
	"goblog-end/config"
	"goblog-end/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.Use(config.Cors()) //开启中间件 允许使用跨域请求

	/**
	* @description:
	* @return {*}
	 */
	r.GET("/page/category/", controller.Category)

	/**
	* @description:根据用户userName获取用户信息
	* @return {*}
	 */
	r.GET("/page/userinfo/:name", controller.PageUserInfo)

	/**
	* @description: 返回文章系列
	* @return {*}
	 */
	r.GET("/post", controller.Post)
	r.GET("/post/:pid", controller.PostByPid)
	// r.GET("/post/uid", controller.PostByUid)

	r.POST("/login/", controller.Login)

	r.POST("/writing", controller.Writing)

	r.POST("/api/upload/img", controller.Upload)
	r.GET("/api/qiniu", controller.QiniuToken)

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
