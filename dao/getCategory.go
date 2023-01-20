/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2022-12-21 15:53:59
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-20 12:55:31
 * @FilePath: /web/goblog-back/dao/getCategory.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	// "goblog-end/model"
	"goblog-end/model/DBmodel"
)

func GetCategory() []model.BlogCategory {
	// fmt.Println("db ==", DB)
	var dataList []model.BlogCategory

	// 查询数据库
	DB.Find(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}
