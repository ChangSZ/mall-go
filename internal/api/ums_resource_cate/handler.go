package ums_resource_cate

import (
	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// ListAll 查询所有后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/listAll [get]
	ListAll() core.HandlerFunc

	// Create 添加后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/create [post]
	Create() core.HandlerFunc

	// Update 修改后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/update/{id} [post]
	Update() core.HandlerFunc

	// Delete 根据ID删除后台资源分类
	// @Tags UmsResourceController
	// @Router /resourceCategory/delete/{id} [post]
	Delete() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	menuService menu.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		menuService: menu.New(db, cache),
	}
}

func (h *handler) i() {}
