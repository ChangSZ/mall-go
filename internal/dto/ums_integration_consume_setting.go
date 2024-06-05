package dto

type UmsIntegrationConsumeSetting struct {
	Id                 int64 `json:"id"`                 //
	DeductionPerAmount int32 `json:"deductionPerAmount"` // 每一元需要抵扣的积分数量
	MaxPercentPerOrder int32 `json:"maxPercentPerOrder"` // 每笔订单最高抵用百分比
	UseUnit            int32 `json:"useUnit"`            // 每次使用积分最小单位100
	CouponStatus       int32 `json:"couponStatus"`       // 是否可以和优惠券同用；0->不可以；1->可以
}
