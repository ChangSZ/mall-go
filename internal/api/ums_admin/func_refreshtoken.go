package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type refreshTokenRequest struct{}

type refreshTokenResponse struct{}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 刷新token
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body refreshTokenRequest true "请求信息"
// @Success 200 {object} refreshTokenResponse
// @Failure 400 {object} code.Failure
// @Router /admin/refreshToken [get]
func (h *handler) RefreshToken() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
