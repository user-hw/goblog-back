package dao

import (
	"fmt"
	"goblog-end/model"
)

func GetUserInfoByName(name string) []model.UserInfo {
	var dataList []model.UserInfo

	// 查询数据库
	DB.Where("user_name = ?", name).First(&dataList)
	fmt.Printf("dataList: %v\n", dataList)

	return dataList
}
