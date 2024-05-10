package ums_member_level

import (
	"github.com/ChangSZ/mall-go/internal/services/ums_member_level"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 查询所有会员等级
	// @Tags UmsMemberLevelController
	// @Router /memberLevel/list [get]
	List(*gin.Context)
}

type handler struct {
	service ums_member_level.Service
}

func New() Handler {
	return &handler{
		service: ums_member_level.New(),
	}
}

func (h *handler) i() {}
