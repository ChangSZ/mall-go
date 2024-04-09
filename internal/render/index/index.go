package index

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
	db     mysql.Repo
}

func New(logger *zap.Logger, db mysql.Repo) *handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) Index() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("index", nil)
	}
}
