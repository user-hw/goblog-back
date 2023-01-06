package controller

import (
	"fmt"
	"goblog-end/dao"
	"goblog-end/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	// 获取路径参数
	// name := c.Param("name")

	var postList []model.BlogPost

	//查询数据库
	postList = dao.GetAllPost()
	for i := 0; i < len(postList); i++ {
		if len(postList[i].Content) > 100 {
			postList[i].Content = postList[i].Content[0:100] + "..."
		}
		fmt.Printf("v: %v\n", postList[i].Content)
	}

	type Res struct {
		Post     model.BlogPost `json:"post"`
		UserInfo model.UserInfo `json:"userInfo"`
	}
	//判断是否查到数据
	if len(postList) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": postList,
		})
	}
}

func PostByPid(c *gin.Context) {
	// 获取路径参数
	pid := c.Param("pid")

	var post []model.BlogPost

	//查询数据库
	post = dao.GetPostByPid(pid)

	uid := post[0].UserId
	fmt.Printf("uid: %v\n", reflect.TypeOf(uid))
	userinfo := dao.GetUserInfoByUid(uid)
	fmt.Printf("userinfo: %v\n", userinfo)

	type Res struct {
		Post     model.BlogPost `json:"post"`
		UserInfo model.UserInfo `json:"userInfo"`
	}
	res := Res{
		Post:     post[0],
		UserInfo: userinfo[0],
	}
	// fmt.Printf("post: %v\n", post[0])
	// for i := 0; i < len(dataList); i++ {
	// 	if len(dataList[i].Content) > 100 {
	// 		dataList[i].Content = dataList[i].Content[0:100] + "..."
	// 	}
	// 	fmt.Printf("v: %v\n", dataList[i].Content)

	// }

	//判断是否查到数据
	if len(post) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": res,
		})
	}

}
