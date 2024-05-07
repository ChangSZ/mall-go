package sms_flash_promotion_product_relation

var _ Service = (*service)(nil)

// 限时购商品关联管理Service
type Service interface {
	i()

	/**
	 * 批量添加关联
	 */

	/**
	 * 修改关联信息
	 */

	/**
	 * 删除关联
	 */

	/**
	 * 获取关联详情
	 */

	/**
	 * 分页查询相关商品及限时购促销信息
	 *
	 * @param flashPromotionId        限时购id
	 * @param flashPromotionSessionId 限时购场次id
	 */

	/**
	 * 根据活动和场次id获取商品关系数量
	 * @param flashPromotionId        限时购id
	 * @param flashPromotionSessionId 限时购场次id
	 */
}
