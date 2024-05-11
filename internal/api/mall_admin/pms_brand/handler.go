package pms_brand

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/pms_brand"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加品牌
	// @Tags PmsBrandController
	// @Router /brand/create [post]
	Create(*gin.Context)

	// Update 更新品牌
	// @Tags PmsBrandController
	// @Router /brand/update/{id} [post]
	Update(*gin.Context)

	// Delete 删除品牌
	// @Tags PmsBrandController
	// @Router /brand/delete/{id} [post]
	Delete(*gin.Context)

	// DeleteBatch 批量删除品牌
	// @Tags PmsBrandController
	// @Router /brand/delete/batch [post]
	DeleteBatch(*gin.Context)

	// List 根据品牌名称分页获取品牌列表
	// @Tags PmsBrandController
	// @Router /brand/list [get]
	List(*gin.Context)

	// ListAll 获取全部品牌列表
	// @Tags PmsBrandController
	// @Router /brand/listAll [get]
	ListAll(*gin.Context)

	// GetItem 根据编号查询品牌信息
	// @Tags PmsBrandController
	// @Router /brand/{id} [get]
	GetItem(*gin.Context)

	// UpdateShowStatus 批量更新显示状态
	// @Tags PmsBrandController
	// @Router /brand/update/showStatus [post]
	UpdateShowStatus(*gin.Context)

	// UpdateFactoryStatus 批量更新厂家制造商状态
	// @Tags PmsBrandController
	// @Router /brand/update/factoryStatus [post]
	UpdateFactoryStatus(*gin.Context)
}

type handler struct {
	service pms_brand.Service
}

func New() Handler {
	return &handler{
		service: pms_brand.New(),
	}
}

func (h *handler) i() {}
