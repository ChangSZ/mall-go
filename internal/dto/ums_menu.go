package dto

import "time"

type UmsMenu struct {
	Id         int64     `json:"id"`
	ParentId   int64     `json:"parentId"`
	CreateTime time.Time `json:"createTime"`
	Title      string    `json:"title"`
	Level      int32     `json:"level"`
	Sort       int32     `json:"sort"`
	Name       string    `json:"name"`
	Icon       string    `json:"icon"`
	Hidden     int32     `json:"hidden"`
}

type UmsMenuParam struct {
	ParentId int64  `json:"parentId"`
	Title    string `json:"title"`
	Level    int32  `json:"level"`
	Sort     int32  `json:"sort"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Hidden   int32  `json:"hidden"`
}

type UmsMenuNode struct {
	Id         int64         `json:"id"`
	ParentId   int64         `json:"parentId"`
	CreateTime time.Time     `json:"createTime"`
	Title      string        `json:"title"`
	Level      int32         `json:"level"`
	Sort       int32         `json:"sort"`
	Name       string        `json:"name"`
	Icon       string        `json:"icon"`
	Hidden     int32         `json:"hidden"`
	Children   []UmsMenuNode `json:"children"`
}

type UmsMenuListUri struct {
	ParentId int64 `uri:"parentId" binding:"required"`
}
