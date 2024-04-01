package ums_integration_consume_setting

// UmsIntegrationConsumeSetting 积分消费设置
//
//go:generate gormgen -structs UmsIntegrationConsumeSetting -input .
type UmsIntegrationConsumeSetting struct {
	Id                 int64 //
	DeductionPerAmount int32 // 每一元需要抵扣的积分数量
	MaxPercentPerOrder int32 // 每笔订单最高抵用百分比
	UseUnit            int32 // 每次使用积分最小单位100
	CouponStatus       int32 // 是否可以和优惠券同用；0->不可以；1->可以
}
