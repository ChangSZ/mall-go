package oms_cart_item

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_cart_item"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Add 添加商品到购物车
	// @Tags OmsCartItemController
	// @Router /cart/add [post]
	Add(*gin.Context)

	// List 获取当前会员的购物车列表
	// @Tags OmsCartItemController
	// @Router /cart/list [get]
	List(*gin.Context)

	// ListPromotion 获取当前会员的购物车列表,包括促销信息
	// @Tags OmsCartItemController
	// @Router /cart/list/promotion [get]
	ListPromotion(*gin.Context)

	// UpdateQuantity 修改购物车中指定商品的数量
	// @Tags OmsCartItemController
	// @Router /cart/update/quantity [get]
	UpdateQuantity(*gin.Context)

	// GetCartProduct 获取购物车中指定商品的规格,用于重选规格
	// @Tags OmsCartItemController
	// @Router /cart/getProduct/{productId} [get]
	GetCartProduct(*gin.Context)

	// UpdateAttr 修改购物车中商品的规格
	// @Tags OmsCartItemController
	// @Router /cart/update/attr [post]
	UpdateAttr(*gin.Context)

	// Delete 删除购物车中的指定商品
	// @Tags OmsCartItemController
	// @Router /cart/delete [post]
	Delete(*gin.Context)

	// Clear 清空当前会员的购物车
	// @Tags OmsCartItemController
	// @Router /cart/clear [post]
	Clear(*gin.Context)
}

type handler struct {
	service oms_cart_item.Service
}

func New() Handler {
	return &handler{
		service: oms_cart_item.New(),
	}
}

func (h *handler) i() {}
