/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-09 22:52:40
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 15:59:01
 * @FilePath: /goblog-back/controller/login.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	// "context"
	"errors"
	"fmt"

	// "net/http"

	"goblog-end/dao"
	"goblog-end/utils"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis"
	// "github.com/go-redis/redis/v8"
	// "github.com/gomodule/redigo/redis"
)

// PingExample godoc
// @Summary 账号密码登录
// @Schemes
// @Description 登录
// @Param UserName MD5Passwd
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /login [get]
func Login(c *gin.Context) {

	type UserLogin struct {
		UserName  string `json:userName`
		MD5Passwd string
	}

	userLogin := UserLogin{}
	c.BindJSON(&userLogin)

	res := dao.GetLogin(userLogin.UserName, userLogin.MD5Passwd)
	if len(res) != 0 {
		uid := res[0].Uid
		// 生成token jwt技术进行生成 A.B.C 生成加密规则 加载数据,过期时间 进行加密算法
		token, err := utils.Award(&uid)
		if err != nil {
			errors.New("token 未能生成")
		} else {
			fmt.Printf("token: %v\n", token)
		}

		// // 将 JWT 存入 Redis 缓存
		dao.TokenToRedis(userLogin.UserName, token)
		// redisClient := Pool.Get()
		// defer redisClient.Close()
		// _, err = redisClient.Do("SET", userLogin.UserName, token, time.Hour*120)
		// if err != nil {
		// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token"})
		// 	// return
		// 	fmt.Printf("err: %v\n", err)
		// }

		userInfo := dao.GetUserInfoByUid(uid)
		// fmt.Printf("userInfo: %v\n", userInfo)

		c.JSON(200, gin.H{
			"meg":      "登录成功",
			"code":     200,
			"token":    token,
			"userInfo": userInfo,
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "登录失败",
			"code": 403,
		})
	}
}
