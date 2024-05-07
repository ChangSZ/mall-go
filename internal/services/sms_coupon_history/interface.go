package sms_coupon_history

var _ Service = (*service)(nil)

// 优惠券领取记录管理Service
type Service interface {
	i()

	/**
	 * 分页查询优惠券领取记录
	 * @param couponId 优惠券id
	 * @param useStatus 使用状态
	 * @param orderSn 使用订单号码
	 */
}
