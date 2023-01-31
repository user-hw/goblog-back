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
	blogPost := model.BlogPost{
		Pid:     pid,
		Title:   title,
		Content: content,
		Uid:     uid,
	}
	_, err := dao.SaveAndUpdatePost(blogPost)
	if err != nil {
		c.JSON(200, gin.H{
			"meg":  "写入失败",
			"code": 400,
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "写入成功",
			"code": 200,
		})
	}

}
