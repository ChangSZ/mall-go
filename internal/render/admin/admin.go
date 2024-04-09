package admin

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

func (h *handler) Login() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("admin_login", nil)
	}
}

func (h *handler) Add() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("admin_add", nil)
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("admin_list", nil)
	}
}

func (h *handler) Menu() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("menu_view", nil)
	}
}

func (h *handler) AdminMenu() core.HandlerFunc {
	type adminMenuRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type adminMenuResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	return func(ctx core.Context) {
		req := new(adminMenuRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		obj := new(adminMenuResponse)
		obj.HashID = req.Id

		ctx.HTML("admin_menu", obj)
	}
}

func (h *handler) MenuAction() core.HandlerFunc {
	type menuActionRequest struct {
		Id string `uri:"id"` // 主键ID
	}

	type menuActionResponse struct {
		HashID string `json:"hash_id"` // hashID
	}

	return func(ctx core.Context) {
		req := new(menuActionRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		obj := new(menuActionResponse)
		obj.HashID = req.Id

		ctx.HTML("menu_action", obj)
	}
}

func (h *handler) ModifyInfo() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("admin_modify_info", nil)
	}
}

func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("admin_modify_password", nil)
	}
}
