/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-04 16:35:31
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-16 15:19:49
 * @FilePath: /goblog-back/controller/post.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"fmt"
	"goblog-end/dao"
	"goblog-end/model/DBmodel"
	"reflect"

	"github.com/gin-gonic/gin"
)

func PostByPid(c *gin.Context) {
	// 获取路径参数
	pid := c.Param("pid")

	var post []model.BlogPost

	//查询数据库
	post = dao.GetPostByPid(pid)

	uid := post[0].Uid
	fmt.Printf("uid: %v\n", reflect.TypeOf(uid))
	userinfo := dao.GetUserInfoByUid(uid)
	fmt.Printf("userinfo: %v\n", userinfo)

	//判断是否查到数据
	if len(post) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		type Res struct {
			Post     model.BlogPost `json:"post"`
			UserInfo model.UserInfo `json:"userInfo"`
		}
		res := Res{
			Post: post[0],
			// UserInfo: userinfo[0],
		}
		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": res,
		})
	}

}
