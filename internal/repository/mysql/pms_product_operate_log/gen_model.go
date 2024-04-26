package pms_product_operate_log

import "time"

// PmsProductOperateLog
//
//go:generate gormgen -structs PmsProductOperateLog -input .
type PmsProductOperateLog struct {
	Id               int64     //
	ProductId        int64     //
	PriceOld         float64   //
	PriceNew         float64   //
	SalePriceOld     float64   //
	SalePriceNew     float64   //
	GiftPointOld     int32     // 赠送的积分
	GiftPointNew     int32     //
	UsePointLimitOld int32     //
	UsePointLimitNew int32     //
	OperateMan       string    // 操作人
	CreateTime       time.Time `gorm:"autoCreateTime"` //
}
