package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type loginRequest struct{}

type loginResponse struct{}

// Login 登录以后返回token
// @Summary 登录以后返回token
// @Description 登录以后返回token
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body loginRequest true "请求信息"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /admin/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
