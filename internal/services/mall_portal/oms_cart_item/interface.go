package oms_cart_item

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 购物车管理Service
type Service interface {
	i()

	/**
	 * 查询购物车中是否包含该商品，有增加数量，无添加到购物车
	 */
	Add(ctx context.Context, param dto.OmsCartItem) (int64, error)

	/**
	 * 根据会员编号获取购物车列表
	 */
	List(ctx context.Context) ([]dto.OmsCartItem, error)

	/**
	 * 获取包含促销活动信息的购物车列表
	 */
	ListPromotion(ctx context.Context, cartIds []int64) ([]dto.CartPromotionItem, error)

	/**
	 * 修改某个购物车商品的数量
	 */
	UpdateQuantity(ctx context.Context, id int64, quantity int32) (int64, error)

	/**
	 * 批量删除购物车中的商品
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 *获取购物车中用于选择商品规格的商品信息
	 */
	GetCartProduct(ctx context.Context, productId int64) (*dto.CartProduct, error)

	/**
	 * 修改购物车中商品的规格
	 */
	UpdateAttr(ctx context.Context, param dto.OmsCartItem) (int64, error)

	/**
	 * 清空购物车
	 */
	Clear(ctx context.Context) (int64, error)
}
