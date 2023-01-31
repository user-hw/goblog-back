/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-31 15:58:46
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 20:33:55
 * @FilePath: /goblog-back/controller/writing.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	// "fmt"

	// "goblog-end/dao"

	_ "fmt"
	"goblog-end/dao"
	model "goblog-end/model/DBmodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Writing(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	title := c.Query("title")
	content := c.Query("content")
	uid, _ := strconv.Atoi(c.Query("uid"))
	text := c.Query("text")
	blogPost := model.BlogPost{
		Pid:     pid,
		Title:   title,
		Content: content,
		Uid:     uid,
		Text:    text,
	}
	_, err := dao.SaveAndUpdatePost(blogPost)
	if err != nil {
		c.JSON(200, gin.H{
			"msg":  "写入失败",
			"code": 400,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "写入成功",
			"code": 200,
		})
	}

}
