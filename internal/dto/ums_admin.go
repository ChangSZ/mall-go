package dto

import "time"

type UmsAdminParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Icon     string `json:"icon"`
	Email    string `json:"email" binding:"email"`
	NickName string `json:"nickName"`
	Note     string `json:"note"`
}

type UmsAdmin struct {
	Id         int64     `json:"id"`
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

type UmsAdminIdUri struct {
	AdminId int64 `uri:"adminId" binding:"required"`
}
