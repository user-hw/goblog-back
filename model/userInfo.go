package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Uid      int    `gorm:"not null;autoIncrement;unique;" json:"uid"`
	UserName string `gorm:"type:varchar(20);not null; default:aaa" json:"userName" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `gorm:"type:varchar(40);" binding:"required" json:"email"`
	Phone    string `gorm:"type:varchar(20);" binding:"required" json:"phone"`
	UserDesc string `gorm:"type:varchar(150);" binding:"required" json:"userDesc"`
	Bilibili string `gorm:"type:varchar(150);" binding:"required" json:"bilibili"`
	Zhihu    string `gorm:"type:varchar(150);" binding:"required" json:"zhihu"`
	Github   string `gorm:"type:varchar(150);" binding:"required" json:"github"`
}
