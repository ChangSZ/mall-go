package pms_product_attr_cate

import (
	"github.com/ChangSZ/mall-go/internal/services/pms_product_attr_cate"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加商品属性分类
	// @Tags PmsProductAttributeCategoryController
	// @Router /productAttribute/category/create [post]
	Create(*gin.Context)

	// Update 修改商品属性分类
	// @Tags PmsProductAttributeCategoryController
	// @Router /productAttribute/category/update/{id} [post]
	Update(*gin.Context)

	// Delete 删除单个商品属性分类
	// @Tags PmsProductAttributeCategoryController
	// @Router /productAttribute/category/delete/{id} [post]
	Delete(*gin.Context)

	// GetItem 获取单个商品属性分类信息
	// @Tags PmsProductAttributeCategoryController
	// @Router /productAttribute/category/{id} [get]
	GetItem(*gin.Context)

	// List 分页获取所有商品属性分类
	// @Tags PmsProductAttributeCategoryController
	// @Router /productAttribute/category/list [get]
	List(*gin.Context)

	// GetListWithAttr 获取所有商品属性分类及其下属性
	// @Tags PmsProductAttributeCategoryController
	// @Router /productAttribute/category/list/withAttr [get]
	GetListWithAttr(*gin.Context)
}

type handler struct {
	service pms_product_attr_cate.Service
}

func New() Handler {
	return &handler{
		service: pms_product_attr_cate.New(),
	}
}

func (h *handler) i() {}
