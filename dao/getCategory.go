package dao

import (
	"goblog-end/model"
)

func GetCategory() []model.BlogCategory {
	// fmt.Println("db ==", DB)
	var dataList []model.BlogCategory

	// 查询数据库
	DB.Find(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}
