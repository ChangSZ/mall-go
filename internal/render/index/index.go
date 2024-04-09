package index

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"

	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{logger: logger}
}

func (h *handler) Index() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("index", nil)
	}
}
