/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2022-12-21 16:55:29
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 13:50:24
 * @FilePath: /goblog-back/model/DBmodel/userInfo.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Uid        int    ` json:"uid"`
	UserName   string `gorm:"type:varchar(20);not null; default:aaa" json:"userName" binding:"required"`
	Avatar     string `json:"avatar"`
	NickName   string `gorm:"type:varchar(20);" json:nickName`
	Occupation string `gorm:"type:varchar(40);" json:occupation`
	Email      string `gorm:"type:varchar(40);" binding:"required" json:"email"`
	Phone      string `gorm:"type:varchar(20);" binding:"required" json:"phone"`
	UserDesc   string `gorm:"type:varchar(150);" binding:"required" json:"userDesc"`
	Bilibili   string `gorm:"type:varchar(150);" binding:"required" json:"bilibili"`
	Zhihu      string `gorm:"type:varchar(150);" binding:"required" json:"zhihu"`
	Github     string `gorm:"type:varchar(150);" binding:"required" json:"github"`
}
