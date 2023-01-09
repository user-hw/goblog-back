package dao

import "goblog-end/model"

func GetLogin(username string, passwd string) []model.BlogUser {
	// fmt.Println("db ==", DB)
	var dataList []model.BlogUser

	// 查询数据库
	DB.Where("userName = ? AND passwd = ?", username, passwd).First(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}
