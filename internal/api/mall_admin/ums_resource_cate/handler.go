package ums_resource_cate

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/services/mall_admin/ums_resource_cate"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// ListAll 查询所有后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/listAll [get]
	ListAll(*gin.Context)

	// Create 添加后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/create [post]
	Create(*gin.Context)

	// Update 修改后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/update/{id} [post]
	Update(*gin.Context)

	// Delete 根据ID删除后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/delete/{id} [post]
	Delete(*gin.Context)
}

type handler struct {
	service ums_resource_cate.Service
}

func New() Handler {
	return &handler{
		service: ums_resource_cate.New(),
	}
}

func (h *handler) i() {}
