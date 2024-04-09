package authorized

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/ChangSZ/mall-go/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增调用方
	// @Tags API.authorized
	// @Router /api/authorized [post]
	Create() core.HandlerFunc

	// CreateAPI 授权调用方接口地址
	// @Tags API.authorized
	// @Router /api/authorized_api [post]
	CreateAPI() core.HandlerFunc

	// List 调用方列表
	// @Tags API.authorized
	// @Router /api/authorized [get]
	List() core.HandlerFunc

	// ListAPI 调用方接口地址列表
	// @Tags API.authorized
	// @Router /api/authorized_api [get]
	ListAPI() core.HandlerFunc

	// Delete 删除调用方
	// @Tags API.authorized
	// @Router /api/authorized/{id} [delete]
	Delete() core.HandlerFunc

	// DeleteAPI 删除调用方接口地址
	// @Tags API.authorized
	// @Router /api/authorized_api/{id} [delete]
	DeleteAPI() core.HandlerFunc

	// UpdateUsed 更新调用方为启用/禁用
	// @Tags API.authorized
	// @Router /api/authorized/used [patch]
	UpdateUsed() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	authorizedService authorized.Service
	hashids           hash.Hash
}

func New(logger *zap.Logger, db mysql.Repo) Handler {
	return &handler{
		logger:            logger,
		authorizedService: authorized.New(db),
		hashids:           hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
