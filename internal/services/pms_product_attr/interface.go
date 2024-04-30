package pms_product_attr

var _ Service = (*service)(nil)

// 商品属性管理Service
type Service interface {
	i()

	/**
	 * 根据分类ID和类型分页获取商品属性
	 * @param cid 分类id
	 * @param type 0->规格；1->参数
	 */

	/**
	 * 添加商品属性
	 */

	/**
	 * 修改商品属性
	 */

	/**
	 * 获取单个商品属性信息
	 */

	/**
	 * 批量删除商品属性
	 */

	/**
	 * 获取商品分类对应属性列表
	 */
}
