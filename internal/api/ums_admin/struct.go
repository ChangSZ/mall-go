package ums_admin

import "time"

type UmsAdmin struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Icon       string    `json:"icon"`
	Email      string    `json:"email"`
	NickName   string    `json:"nickName"`
	Note       string    `json:"note"`
	CreateTime time.Time `json:"createTime"`
	LoginTime  time.Time `json:"loginTime"`
	Status     int32     `json:"status"`
}
