package pms_product

var _ Service = (*service)(nil)

// 商品管理Service
type Service interface {
	i()

	/**
	 * 创建商品
	 */

	/**
	 * 根据商品ID获取商品信息（用于更新商品）
	 */

	/**
	 * 更新商品
	 */

	/**
	 * 分页查询商品
	 */

	/**
	 * 批量修改审核状态
	 * @param ids 商品ID列表
	 * @param verifyStatus 审核状态
	 * @param detail 审核详情
	 */

	/**
	 * 批量修改商品上架状态
	 */

	/**
	 * 批量修改商品推荐状态
	 */

	/**
	 * 批量修改新品状态
	 */

	/**
	 * 批量删除商品
	 */

	/**
	 * 根据商品名称或者货号模糊查询
	 */
}
