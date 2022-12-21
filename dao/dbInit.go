package dao

import (
	"goblog-end/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {

	dsn := "root:123456@tcp(127.0.0.1)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("aaaaaaaa")
	}

	db.AutoMigrate(&model.UserInfo{})
	db.AutoMigrate(&model.BlogCategory{})
	DB = db

}
