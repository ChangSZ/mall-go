package authorized

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/ChangSZ/mall-go/pkg/hash"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增调用方
	// @Tags API.authorized
	// @Router /api/authorized [post]
	Create(*gin.Context)

	// CreateAPI 授权调用方接口地址
	// @Tags API.authorized
	// @Router /api/authorized_api [post]
	CreateAPI(*gin.Context)

	// List 调用方列表
	// @Tags API.authorized
	// @Router /api/authorized [get]
	List(*gin.Context)

	// ListAPI 调用方接口地址列表
	// @Tags API.authorized
	// @Router /api/authorized_api [get]
	ListAPI(*gin.Context)

	// Delete 删除调用方
	// @Tags API.authorized
	// @Router /api/authorized/{id} [delete]
	Delete(*gin.Context)

	// DeleteAPI 删除调用方接口地址
	// @Tags API.authorized
	// @Router /api/authorized_api/{id} [delete]
	DeleteAPI(*gin.Context)

	// UpdateUsed 更新调用方为启用/禁用
	// @Tags API.authorized
	// @Router /api/authorized/used [patch]
	UpdateUsed(*gin.Context)
}

type handler struct {
	service authorized.Service
	hashids hash.Hash
}

func New() Handler {
	return &handler{
		service: authorized.New(),
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
