package pms_portal_product

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/pms_portal_product"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Search 综合搜索、筛选、排序
	// @Tags PmsPortalProductController
	// @Router /product/search [get]
	Search(*gin.Context)

	// CategoryTreeList 以树形结构获取所有商品分类
	// @Tags PmsPortalProductController
	// @Router /product/categoryTreeList [get]
	CategoryTreeList(*gin.Context)

	// Detail 获取前台商品详情
	// @Tags PmsPortalProductController
	// @Router /product/detail/{id} [get]
	Detail(*gin.Context)
}

type handler struct {
	service pms_portal_product.Service
}

func New() Handler {
	return &handler{
		service: pms_portal_product.New(),
	}
}

func (h *handler) i() {}
