package dto

import "time"

type OmsOrderReturnReason struct {
	Id         int64     `json:"id"`         //
	Name       string    `json:"name"`       // 退货类型
	Sort       int32     `json:"sort"`       //
	Status     int32     `json:"status"`     // 状态：0->不启用；1->启用
	CreateTime time.Time `json:"createTime"` // 添加时间
}
