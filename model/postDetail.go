package model

/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-20 12:23:22
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-21 22:20:29
 * @FilePath: /web/goblog-back/model/postDetail.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"goblog-end/model/DBmodel"
)

type PostDetail struct {
	BlogPost model.BlogPost ` json:"blogPost"`
	// model.BlogPost
	Uid       int    ` json:"uid"`
	UserName  string `gorm:"type:varchar(20);not null; default:aaa" json:"userName" binding:"required"`
	Avatar    string `json:"avatar"`
	NickName  string `gorm:"type:varchar(20);" json:nickName`
	PageCount int    `json:"pageCount"`
}
