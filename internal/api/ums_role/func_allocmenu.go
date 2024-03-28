package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type allocMenuRequest struct{}

type allocMenuResponse struct{}

// AllocMenu 给角色分配菜单
// @Summary 给角色分配菜单
// @Description 给角色分配菜单
// @Tags ROLE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body allocMenuRequest true "请求信息"
// @Success 200 {object} allocMenuResponse
// @Failure 400 {object} code.Failure
// @Router /role/allocMenu [post]
func (h *handler) AllocMenu() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
