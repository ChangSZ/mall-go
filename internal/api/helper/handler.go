package helper

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/services/authorized"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Md5 加密
	// @Tags Helper
	// @Router /helper/md5/{str} [get]
	Md5() core.HandlerFunc

	// Sign 签名
	// @Tags Helper
	// @Router /helper/sign [post]
	Sign() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	db                mysql.Repo
	authorizedService authorized.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:            logger,
		db:                db,
		authorizedService: authorized.New(db, cache),
	}
}

func (h *handler) i() {}
