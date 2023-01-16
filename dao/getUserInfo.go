/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2022-12-21 16:56:25
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-16 15:30:17
 * @FilePath: /goblog-back/dao/getUserInfo.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
func GetUserInfoByUid(uid int) []model.UserInfo {
	var dataList []model.UserInfo

	// 查询数据库
	DB.Where("uid= ?", uid).First(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)
	if len(dataList) == 0 {
		return nil
	} else {

		return dataList
	}

}
