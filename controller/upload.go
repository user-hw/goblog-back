/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-31 21:12:21
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 21:27:37
 * @FilePath: /goblog-back/controller/upload.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	image := c.PostForm("wangeditor-uploaded-image")
	fmt.Printf("image: %v\n", image)
	c.JSON(200, gin.H{
		"meg":  "图片成功",
		"code": 200,
	})
}
