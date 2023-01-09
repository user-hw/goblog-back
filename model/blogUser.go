package model

import "time"

type BlogUser struct {
	Uid      int       `json:"uid"`
	UserName string    `json:"userName"`
	Passwd   string    `json:"passwd"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
