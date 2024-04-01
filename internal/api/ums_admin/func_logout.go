package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type logoutRequest struct{}

type logoutResponse struct{}

// Logout 登出功能
// @Summary 登出功能
// @Description 登出功能
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body logoutRequest true "请求信息"
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /admin/logout [post]
func (h *handler) Logout() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
