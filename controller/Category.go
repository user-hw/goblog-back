package controller

import (
	"goblog-end/dao"
	"goblog-end/model"

	"github.com/gin-gonic/gin"
)

func Category(c *gin.Context) {

	// 获取路径参数
	// name := c.Param("name")

	var dataList []model.BlogCategory

	//查询数据库
	dataList = dao.GetCategory()

	//判断是否查到数据
	if len(dataList) == 0 {
		c.JSON(200, gin.H{
			"meg":  "没有查到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"meg":  "查询成功",
			"code": 200,
			"data": dataList,
		})
	}

}
