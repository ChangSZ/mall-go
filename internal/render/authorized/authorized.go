package authorized

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
		ctx.HTML("authorized_add", nil)
	}
}

func (h *handler) Demo() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_demo", nil)
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_list", nil)
	}
}

func (h *handler) Api() core.HandlerFunc {
	type apiRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type apiResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	return func(ctx core.Context) {
		req := new(apiRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		obj := new(apiResponse)
		obj.HashID = req.Id

		ctx.HTML("authorized_api", obj)
	}
}
