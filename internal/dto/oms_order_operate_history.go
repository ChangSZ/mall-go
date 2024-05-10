package dto

import "time"

type OmsOrderOperateHistory struct {
	Id          int64     `json:"id"`          //
	OrderId     int64     `json:"orderId"`     // 订单id
	OperateMan  string    `json:"operateMan"`  // 操作人：用户；系统；后台管理员
	CreateTime  time.Time `json:"createTime"`  // 操作时间
	OrderStatus int32     `json:"orderStatus"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string    `json:"note"`        // 备注
}
