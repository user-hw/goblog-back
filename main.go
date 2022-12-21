package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Products struct {
	gorm.Model
	Code  string
	Price uint
}

type UserInfo struct {
	gorm.Model
	Uid      int    `json:"uid"`
	UserName string `gorm:"type:varchar(20);not null" json:"userName" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `gorm:"type:varchar(40);not null" binding:"required" json:email`
	Phone    string `gorm:"type:varchar(20);not null" binding:"required" json:phone`
}

func create(db *gorm.DB) {
	db.AutoMigrate(&Products{})
}

func insert(db *gorm.DB) {
	p := Products{
		Code:  "1001",
		Price: 100,
	}
	db.Create(&p)
}

func find(db *gorm.DB) {
	var p Products
	db.First(&p, 1)
	fmt.Printf("p: %v\n", p)
	db.First(&p, "Code= ?", "1001")
	fmt.Printf("p: %v\n", p)
}

func update(db *gorm.DB) {
	var p Products
	db.First(&p, 1)
	db.Model(&p).Update("Price", 1000)

}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")                                                     // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //你想放行的header也可以在后面自行添加
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")              //我自己只使用 get post 所以只放行它
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("aaaaaaaa")
	}
	fmt.Println("db", db)
	db.AutoMigrate(&UserInfo{})

	r := gin.Default()
	r.Use(Cors()) //开启中间件 允许使用跨域请求

	r.GET("/", func(c *gin.Context) {
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

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "请求user",
			"msg":  "user你好吗",
		})

	})

	port := "8082"
	r.Run(":" + port)

}
