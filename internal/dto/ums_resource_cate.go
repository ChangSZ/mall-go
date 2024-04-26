package dto

import "time"

type UmsResourceCateParam struct {
	Name string `json:"name"`
	Sort int32  `json:"sort"`
}

type UmsResourceCate struct {
	Id         int64     `json:"id"`
	CreateTime time.Time `json:"createTime"`
	Name       string    `json:"name"`
	Sort       int32     `json:"sort"`
}
