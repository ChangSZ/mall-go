package oms_order_return_reason

import "time"

// OmsOrderReturnReason 退货原因表
//
//go:generate gormgen -structs OmsOrderReturnReason -input .
type OmsOrderReturnReason struct {
	Id         int64     //
	Name       string    // 退货类型
	Sort       int32     //
	Status     int32     // 状态：0->不启用；1->启用
	CreateTime time.Time `gorm:"autoCreateTime"` // 添加时间
}
