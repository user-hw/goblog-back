/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-04 16:35:31
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-02-03 18:01:24
 * @FilePath: /goblog-back/controller/post.go
 * @Description: 获取所有文章列表
 */
package controller

import (
	// "fmt"

	"fmt"
	"goblog-end/dao"
	mymodel "goblog-end/model"
	model "goblog-end/model/DBmodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	// 获取路径参数
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	postNum, _ := strconv.Atoi(c.Query("postNum"))
	uid, _ := strconv.Atoi(c.Query("uid"))
	// fmt.Printf("id: %v\n", id)

	var postList []model.BlogPost

	fmt.Printf("uid: %v\n", uid)

	//查询数据库
	if uid != 0 {
		postList = dao.GetPostByUid(uid)
	} else {
		postList = dao.GetAllPost()
	}
	for i := 0; i < len(postList); i++ {
		text := postList[i].Text
		// go输入中文是3个字节为一个单位，截取的时候要使用3的倍数或者采用下面的方法转化
		// nameRune := []rune(text)

		if len(text) > 99 {
			fmt.Printf("postList[i].Text: %v\n", text[0:99])
			postList[i].Text = text[0:99] + "..."
		}
	}

	//判断是否查到数据
	if len(postList) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		postAll := []mymodel.PostDetail{}
		len_postList := len(postList)
		startNum := (pageNum - 1) * postNum
		endNum := 0

		if startNum+postNum > len_postList {
			endNum = len_postList
		} else {
			endNum = startNum + postNum
		}
		fmt.Printf("startNum: %v\n", startNum)
		fmt.Printf("endNum: %v\n", endNum)

		for i := startNum; i < endNum; i++ {
			uid := postList[i].Uid
			userInfo := dao.GetUserInfoByUid(uid)
			post := mymodel.PostDetail{
				BlogPost:  postList[i],
				Uid:       userInfo[0].Uid,
				UserName:  userInfo[0].UserName,
				NickName:  userInfo[0].NickName,
				PageCount: (len(postList)-1)/postNum + 1,
			}
			postAll = append(postAll, post)

		}
		// fmt.Printf("postAll: %v\n", postAll)

		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": postAll,
		})
	}
}
