package pms_product_cate

import (
	"github.com/ChangSZ/mall-go/internal/services/pms_product_cate"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加商品分类
	// @Tags PmsProductCategoryController
	// @Router /productCategory/create [post]
	Create(*gin.Context)

	// Update 修改商品分类
	// @Tags PmsProductCategoryController
	// @Router /productCategory/update/{id} [post]
	Update(*gin.Context)

	// List 分页查询商品分类
	// @Tags PmsProductCategoryController
	// @Router /productCategory/{parentId} [get]
	List(*gin.Context)

	// GetItem 根据id获取商品分类
	// @Tags PmsProductCategoryController
	// @Router /productCategory/{id} [get]
	GetItem(*gin.Context)

	// Delete 删除商品分类
	// @Tags PmsProductCategoryController
	// @Router /productCategory/delete/{id} [post]
	Delete(*gin.Context)

	// UpdateNavStatus 修改导航栏显示状态
	// @Tags PmsProductCategoryController
	// @Router /productCategory/update/navStatus[post]
	UpdateNavStatus(*gin.Context)

	// UpdateShowStatus 修改显示状态
	// @Tags PmsProductCategoryController
	// @Router /productCategory/update/showStatus[post]
	UpdateShowStatus(*gin.Context)

	// ListWithChildren 查询所有一级分类及子分类
	// @Tags PmsProductCategoryController
	// @Router /productCategory/list/withChildren [get]
	ListWithChildren(*gin.Context)
}

type handler struct {
	pmsProductCateService pms_product_cate.Service
}

func New() Handler {
	return &handler{
		pmsProductCateService: pms_product_cate.New(),
	}
}

func (h *handler) i() {}
