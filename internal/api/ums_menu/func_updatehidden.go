package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updateHiddenRequest struct{}

type updateHiddenResponse struct{}

// UpdateHidden 修改菜单显示状态
// @Summary 修改菜单显示状态
// @Description 修改菜单显示状态
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateHiddenRequest true "请求信息"
// @Success 200 {object} updateHiddenResponse
// @Failure 400 {object} code.Failure
// @Router /menu/updateHidden/{id} [post]
func (h *handler) UpdateHidden() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
