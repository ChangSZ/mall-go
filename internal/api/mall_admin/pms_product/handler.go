package pms_product

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/mall_admin/pms_product"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建商品
	// @Tags PmsProductController
	// @Router /product/create [post]
	Create(*gin.Context)

	// GetUpdateInfo 根据商品id获取商品编辑信息
	// @Tags PmsProductController
	// @Router /product/updateInfo/{id} [get]
	GetUpdateInfo(*gin.Context)

	// Update 更新商品
	// @Tags PmsProductController
	// @Router /product/update/{id} [post]
	Update(*gin.Context)

	// List 查询商品
	// @Tags PmsProductController
	// @Router /product/list [get]
	List(*gin.Context)

	// SimpleList 根据商品名称或货号模糊查询
	// @Tags PmsProductController
	// @Router /product/simpleList [get]
	SimpleList(*gin.Context)

	// UpdateVerifyStatus 批量修改审核状态
	// @Tags PmsProductController
	// @Router /product/update/verifyStatus [post]
	UpdateVerifyStatus(*gin.Context)

	// UpdatePublishStatus 批量上下架商品
	// @Tags PmsProductController
	// @Router /product/update/publishStatus [post]
	UpdatePublishStatus(*gin.Context)

	// UpdateRecommendStatus 批量推荐商品
	// @Tags PmsProductController
	// @Router /product/update/recommendStatus [post]
	UpdateRecommendStatus(*gin.Context)

	// UpdateNewStatus 批量设为新品
	// @Tags PmsProductController
	// @Router /product/update/newStatus [post]
	UpdateNewStatus(*gin.Context)

	// UpdateDeleteStatus 批量修改删除状态
	// @Tags PmsProductController
	// @Router /product/update/deleteStatus [post]
	UpdateDeleteStatus(*gin.Context)
}

type handler struct {
	service pms_product.Service
}

func New() Handler {
	return &handler{
		service: pms_product.New(),
	}
}

func (h *handler) i() {}
