/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2022-12-21 15:45:31
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 16:43:20
 * @FilePath: /goblog-back/dao/dbInit.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	// "goblog-end/model"
	model "goblog-end/model/DBmodel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"goblog-end/config"
)

var DB *gorm.DB

func init() {
	password:= config.Cfg.System.DBPassword
	dsn := "root:"+password+"@tcp(127.0.0.1)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
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
	db.AutoMigrate(&model.BlogPost{})
	db.AutoMigrate(&model.BlogPostA{})
	DB = db

}
