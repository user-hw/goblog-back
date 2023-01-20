/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-09 23:51:21
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-20 12:56:03
 * @FilePath: /web/goblog-back/dao/login.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"goblog-end/model/DBmodel"
)

func GetLogin(username string, passwd string) []model.BlogUser {
	// fmt.Println("db ==", DB)
	var dataList []model.BlogUser

	// 查询数据库
	DB.Where("userName = ? AND passwd = ?", username, passwd).First(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}
