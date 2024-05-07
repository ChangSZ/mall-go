package sms_coupon

var _ Service = (*service)(nil)

// 优惠券管理Service
type Service interface {
	i()

	/**
	 * 添加优惠券
	 */

	/**
	 * 根据优惠券id删除优惠券
	 */

	/**
	 * 根据优惠券id更新优惠券信息
	 */

	/**
	 * 分页获取优惠券列表
	 */

	/**
	 * 获取优惠券详情
	 * @param id 优惠券表id
	 */
}
