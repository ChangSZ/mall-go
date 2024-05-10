package pms_product_attr

import (
	"github.com/ChangSZ/mall-go/internal/services/pms_product_attr"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加商品属性信息
	// @Tags PmsProductAttributeController
	// @Router /productAttribute/create [post]
	Create(*gin.Context)

	// Update 修改商品属性信息
	// @Tags PmsProductAttributeController
	// @Router /productAttribute/update/{id} [post]
	Update(*gin.Context)

	// List 根据分类查询属性列表或参数列表
	// @Tags PmsProductAttributeController
	// @Router /productAttribute/list/{cid} [get]
	List(*gin.Context)

	// GetItem 查询单个商品属性
	// @Tags PmsProductAttributeController
	// @Router /productAttribute/{id} [get]
	GetItem(*gin.Context)

	// Delete 批量删除商品属性
	// @Tags PmsProductAttributeController
	// @Router /productAttribute/delete [post]
	Delete(*gin.Context)

	// GetAttrInfo 根据商品分类的id获取商品属性及属性分类
	// @Tags PmsProductAttributeController
	// @Router /productAttribute/attrInfo/{productCategoryId} [get]
	GetAttrInfo(*gin.Context)
}

type handler struct {
	service pms_product_attr.Service
}

func New() Handler {
	return &handler{
		service: pms_product_attr.New(),
	}
}

func (h *handler) i() {}
