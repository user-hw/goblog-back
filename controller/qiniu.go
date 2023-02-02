/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-31 21:54:19
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-02-02 11:01:52
 * @FilePath: /goblog-back/controller/qiniu.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"goblog-end/config"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func QiniuToken(c *gin.Context) {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "img-store-new"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期

	QiniuAccessKey := config.Cfg.System.QiniuAccessKey
	QiniuSecretKey := config.Cfg.System.QiniuSecretKey

	mac := qbox.NewMac(QiniuAccessKey, QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	c.JSON(200, gin.H{
		"msg":     "获取token成功",
		"code":    200,
		"upToken": upToken,
	})
}
