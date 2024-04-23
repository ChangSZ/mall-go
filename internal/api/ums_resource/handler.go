package ums_resource

import (
	"github.com/ChangSZ/mall-go/internal/services/ums_resource"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加后台资源
	// @Tags UmsResourceController
	// @Router /resource/create [post]
	Create(*gin.Context)

	// Update 修改后台资源
	// @Tags UmsResourceController
	// @Router /resource/update/{id} [post]
	Update(*gin.Context)

	// Get 根据ID获取资源详情
	// @Tags UmsResourceController
	// @Router /resource/{id} [get]
	Get(*gin.Context)

	// Delete 根据ID删除后台资源
	// @Tags UmsResourceController
	// @Router /resource/delete/{id} [post]
	Delete(*gin.Context)

	// List 分页模糊查询后台资源
	// @Tags UmsResourceController
	// @Router /resource/list [get]
	List(*gin.Context)

	// ListAll 查询所有后台资源
	// @Tags UmsResourceController
	// @Router /resource/listAll [get]
	ListAll(*gin.Context)
}

type handler struct {
	umsResourceService ums_resource.Service
}

func New() Handler {
	return &handler{
		umsResourceService: ums_resource.New(),
	}
}

func (h *handler) i() {}
