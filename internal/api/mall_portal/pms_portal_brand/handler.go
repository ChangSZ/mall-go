package pms_portal_brand

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/pms_portal_brand"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// RecommendList 分页获取推荐品牌
	// @Tags PmsPortalBrandController
	// @Router /brand/recommendList [get]
	RecommendList(*gin.Context)

	// Detail 获取品牌详情
	// @Tags PmsPortalBrandController
	// @Router /brand/detail/{brandId} [get]
	Detail(*gin.Context)

	// ProductList 分页获取品牌相关商品
	// @Tags PmsPortalBrandController
	// @Router /brand/productList [get]
	ProductList(*gin.Context)
}

type handler struct {
	service pms_portal_brand.Service
}

func New() Handler {
	return &handler{
		service: pms_portal_brand.New(),
	}
}

func (h *handler) i() {}
