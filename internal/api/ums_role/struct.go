package ums_role

import "time"

type UmsRoleParam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AdminCount  int32  `json:"adminCount"`
	Status      int32  `json:"status"`
	Sort        int32  `json:"sort"`
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

type UmsRoleUri struct {
	Id int64 `uri:"id" binding:"required"`
}

type UmsRoleIdUri struct {
	RoleId int64 `uri:"roleId" binding:"required"`
}
