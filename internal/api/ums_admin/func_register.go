package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type registerRequest struct{}

type registerResponse struct{}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body registerRequest true "请求信息"
// @Success 200 {object} registerResponse
// @Failure 400 {object} code.Failure
// @Router /admin/register [post]
func (h *handler) Register() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
