package oms_order_operate_history

import "time"

// OmsOrderOperateHistory 订单操作历史记录
//
//go:generate gormgen -structs OmsOrderOperateHistory -input .
type OmsOrderOperateHistory struct {
	Id          int64     //
	OrderId     int64     // 订单id
	OperateMan  string    // 操作人：用户；系统；后台管理员
	CreateTime  time.Time `gorm:"autoCreateTime"` // 操作时间
	OrderStatus int32     // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string    // 备注
}
