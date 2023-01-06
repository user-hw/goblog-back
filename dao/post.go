package dao

import "goblog-end/model"

func GetAllPost() []model.BlogPost {
	// fmt.Println("db ==", DB)
	var dataList []model.BlogPost

	// 查询数据库
	DB.Find(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}

func GetPostByPid(pid string) []model.BlogPost {
	var dataList []model.BlogPost

	// 查询数据库
	DB.Where("pid = ?", pid).First(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList

}
