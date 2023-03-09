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
	"fmt"
	model "goblog-end/model/DBmodel"
	_ "time"
)

func GetLogin(username string, passwd string) []model.BlogUser {
	// fmt.Println("db ==", DB)
	var dataList []model.BlogUser

	// 查询数据库
	DB.Where("userName = ? AND passwd = ?", username, passwd).First(&dataList)
	// fmt.Printf("dataList: %v\n", dataList)

	return dataList
}

func TokenToRedis(userName string, token string) {
	// 将 JWT 存入 Redis 缓存
	redisClient := Pool.Get()
	// fmt.Printf("redisClient: %v\n", redisClient)
	defer redisClient.Close()
	redisClient.Do("SET", userName, token)
	_, err := redisClient.Do("EXPIRE", userName, 10*1000*60*60)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token"})
		// return
		fmt.Printf("err: %v\n", err)
	}

}
