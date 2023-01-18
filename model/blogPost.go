/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-04 16:32:27
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-18 09:55:54
 * @FilePath: /goblog-back/model/blogPost.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"time"

	_ "gorm.io/gorm"
)

type BlogPost struct {
	// gorm.Model
	Pid          int       `json:"pid"`          // 文章ID
	Title        string    `json:"title"`        // 文章标题
	Slug         string    `json:"slug"`         // 自定也页面 path
	Content      string    `json:"content"`      // 文章的html
	Markdown     string    `json:"markdown"`     // 文章的Markdown
	CategoryId   int       `json:"categoryId"`   // 分类id
	Uid          int       `json:"Uid "`         // 用户id
	ViewCount    int       `json:"viewCount"`    // 查看次数
	CommentCount int       `json:"commentCount"` // 评论次数
	Type         int       `json:"type"`         //文章类型 0 普通，1 自定义文章
	CreateAt     time.Time `json:"createAt"`     // 创建时间
	UpdateAt     time.Time `json:"updateAt"`     // 更新时间

}
