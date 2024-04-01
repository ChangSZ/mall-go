package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type infoRequest struct{}

type infoResponse struct{}

// Info 获取当前登录用户信息
// @Summary 获取当前登录用户信息
// @Description 获取当前登录用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body infoRequest true "请求信息"
// @Success 200 {object} infoResponse
// @Failure 400 {object} code.Failure
// @Router /admin/info [get]
func (h *handler) Info() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
