package main

import (
	_ "goblog-end/dao"
	"goblog-end/router"
)

func main() {
	router.Start()
	// dsn := "root:123456@tcp(127.0.0.1)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })
	// if err != nil {
	// 	panic("aaaaaaaa")
	// }
	// fmt.Println("db", db)
	// // db.AutoMigrate(&UserInfo{})
	// db.AutoMigrate(&BlogCategory{})

	// r := gin.Default()
	// r.Use(Cors()) //开启中间件 允许使用跨域请求

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"Avatar":   "/resource/images/3.jpg",
	// 		"UserName": "不会写代码的建筑师",
	// 		"UserDesc": "写代码是医生的邀请",
	// 		"Bilibili": "https://space.bilibili.com/66628849",
	// 		"zhihu":    "https://www.zhihu.com/people/xu-heng-wei-63",
	// 		"github":   "https://github.com/user-hw",
	// 		"code":     200,
	// 		// "Categorys": [1,2],
	// 	})

	// })

	// r.GET("/admin/userinfo/:name", func(c *gin.Context) {
	// 	// 获取路径参数
	// 	// name := c.Param("name")

	// 	var dataList []BlogCategory

	// 	//查询数据库
	// 	db.Find(&dataList)

	// 	//判断是否查到数据
	// 	if len(dataList) == 0 {
	// 		c.JSON(200, gin.H{
	// 			"meg":  "没有查到数据",
	// 			"code": 400,
	// 			"data": gin.H{},
	// 		})
	// 	} else {
	// 		c.JSON(200, gin.H{
	// 			"meg":  "查询成功",
	// 			"code": 200,
	// 			"data": dataList,
	// 		})
	// 	}

	// })

	// port := "8082"
	// r.Run(":" + port)

}
