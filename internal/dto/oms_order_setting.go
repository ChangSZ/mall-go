package dto

type OmsOrderSetting struct {
	Id                  int64 `json:"id"`                  //
	FlashOrderOvertime  int32 `json:"flashOrderOvertime"`  // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int32 `json:"normalOrderOvertime"` // 正常订单超时时间(分)
	ConfirmOvertime     int32 `json:"confirmOvertime"`     // 发货后自动确认收货时间（天）
	FinishOvertime      int32 `json:"finishOvertime"`      // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int32 `json:"commentOvertime"`     // 订单完成后自动好评时间（天）
}
