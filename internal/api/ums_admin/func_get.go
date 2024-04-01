package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type getRequest struct{}

type getResponse struct{}

// Get 获取指定用户信息
// @Summary 获取指定用户信息
// @Description 获取指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} getResponse
// @Failure 400 {object} code.Failure
// @Router /admin/{id} [get]
func (h *handler) Get() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
