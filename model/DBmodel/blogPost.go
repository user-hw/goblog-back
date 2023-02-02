/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-04 16:32:27
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-02-02 23:33:35
 * @FilePath: /goblog-back/model/blogPost.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	// "time"

	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Pid          int    `json:"pid"`          // 文章ID
	Title        string `json:"title"`        // 文章标题
	Content      string `json:"content"`      // 文章的html
	Text         string `json:"text"`         // 文章的文字
	TitleImg     string `json:"titleImg"`     // 封面图片
	ShortText    string `json:"shortText"`    // 简介文字，留着备用
	CategoryId   int    `json:"categoryId"`   // 分类id
	Uid          int    `json:"uid "`         // 用户id
	ViewCount    int    `json:"viewCount"`    // 查看次数
	CommentCount int    `json:"commentCount"` // 评论次数
	LikeCount    int    `json:"likeCount"`    // 点赞次数
	SubCount     int    `json:"subCount"`     // 订阅次数
	MetaUrl      string `json:"metaUrl"`      //原文链接
	Type         int    `json:"type"`         //文章类型 0 普通，1 自定义文章
}

type BlogPostA struct {
	gorm.Model
	Pid          int    `json:"pid"`          // 文章ID
	Title        string `json:"title"`        // 文章标题
	Slug         string `json:"slug"`         // 自定也页面 path
	Content      string `json:"content"`      // 文章的html
	Markdown     string `json:"markdown"`     // 文章的Markdown
	CategoryId   int    `json:"categoryId"`   // 分类id
	Uid          int    `json:"Uid "`         // 用户id
	ViewCount    int    `json:"viewCount"`    // 查看次数
	CommentCount int    `json:"commentCount"` // 评论次数
	LikeCount    int    `json:"likeCount"`    // 点赞次数
	SubCount     int    `json:"subCount"`     // 订阅次数
	Type         int    `json:"type"`         //文章类型 0 普通，1 自定义文章
	// CreateAt     time.Time `json:"createAt"`     // 创建时间
	// UpdateAt     time.Time `json:"updateAt"`     // 更新时间
}
