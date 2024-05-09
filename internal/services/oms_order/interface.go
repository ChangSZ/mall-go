package oms_order

var _ Service = (*service)(nil)

// 订单管理Service
type Service interface {
	i()

	/**
	 * 分页查询订单
	 */

	/**
	 * 批量发货
	 */

	/**
	 * 批量关闭订单
	 */

	/**
	 * 批量删除订单
	 */

	/**
	 * 获取指定订单详情
	 */

	/**
	 * 修改订单收货人信息
	 */

	/**
	 * 修改订单费用信息
	 */

	/**
	 * 修改订单备注
	 */
}
