package home

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/home"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Content 首页内容信息展示
	// @Tags HomeController
	// @Router /home/content [get]
	Content(*gin.Context)

	// RecommendProductList 分页获取推荐商品
	// @Tags HomeController
	// @Router /home/recommendProductList [get]
	RecommendProductList(*gin.Context)

	// GetProductCateList 获取首页商品分类
	// @Tags HomeController
	// @Router /home/productCateList/{parentId} [get]
	GetProductCateList(*gin.Context)

	// GetSubjectList 根据分类获取专题
	// @Tags HomeController
	// @Router /home/subjectList [get]
	GetSubjectList(*gin.Context)

	// HotProductList 分页获取人气推荐商品
	// @Tags HomeController
	// @Router /home/hotProductList [get]
	HotProductList(*gin.Context)

	// NewProductList 分页获取新品推荐商品
	// @Tags HomeController
	// @Router /home/newProductList [get]
	NewProductList(*gin.Context)
}

type handler struct {
	service home.Service
}

func New() Handler {
	return &handler{
		service: home.New(),
	}
}

func (h *handler) i() {}
