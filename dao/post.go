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

func GetPostByPid() []model.BlogPost {
	var dataList []model.BlogPost

	// 查询数据库
	DB.Find(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList

}
