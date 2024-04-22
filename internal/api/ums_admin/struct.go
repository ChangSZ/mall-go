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

type UmsAdminUri struct {
	Id int64 `uri:"id" binding:"required"` // 用户ID
}

type UmsAdminIdUri struct {
	AdminId int64 `uri:"adminId" binding:"required"`
}

type UmsRole struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AdminCount  int32     `json:"adminCount"`
	CreateTime  time.Time `json:"createTime"`
	Status      int32     `json:"status"`
	Sort        int32     `json:"sort"`
}
