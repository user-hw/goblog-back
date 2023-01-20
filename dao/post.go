/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-04 16:31:33
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-20 13:06:44
 * @FilePath: /web/goblog-back/dao/post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"goblog-end/model/DBmodel"
)

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

func GetPostByUid(uid int) []model.BlogPost {
	var dataList []model.BlogPost

	// 查询数据库
	DB.Where("uid = ?", uid).Find(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}
