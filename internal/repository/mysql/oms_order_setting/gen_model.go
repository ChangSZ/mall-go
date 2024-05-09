package oms_order_setting

// OmsOrderSetting 订单设置表
//
//go:generate gormgen -structs OmsOrderSetting -input .
type OmsOrderSetting struct {
	Id                  int64 //
	FlashOrderOvertime  int32 // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int32 // 正常订单超时时间(分)
	ConfirmOvertime     int32 // 发货后自动确认收货时间（天）
	FinishOvertime      int32 // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int32 // 订单完成后自动好评时间（天）
}
