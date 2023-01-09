package controller

import (
	"errors"
	"fmt"
	"goblog-end/dao"
	"goblog-end/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type UserLogin struct {
		UserName  string `json:userName`
		MD5Passwd string
	}
	type UserInfo struct {
		Uid      int    `json:"uid"`
		UserName string `json:"userName"`
		Avatar   string `json:"avatar"`
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

		var userInfo UserInfo
		userInfo.Uid = res[0].Uid
		userInfo.UserName = res[0].UserName
		userInfo.Avatar = res[0].Avatar

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
