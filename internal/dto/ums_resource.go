package dto

import "time"

type UmsResourceParam struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	CategoryId  int64  `json:"categoryId"`
}

type UmsResource struct {
	Id          int64     `json:"id"`
	CreateTime  time.Time `json:"createTime"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	CategoryId  int64     `json:"categoryId"`
}
