package cron

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/core"

	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *handler {
	return &handler{logger: logger}
}

func (h *handler) Add() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("cron_task_add", nil)
	}
}

func (h *handler) Edit() core.HandlerFunc {
	type editRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type editResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	return func(ctx core.Context) {
		req := new(editRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		obj := new(editResponse)
		obj.HashID = req.Id
		ctx.HTML("cron_task_edit", obj)
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("cron_task_list", nil)
	}
}
