package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updatePasswordRequest struct{}

type updatePasswordResponse struct{}

// UpdatePassword 修改指定用户密码
// @Summary 修改指定用户密码
// @Description 修改指定用户密码
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePasswordRequest true "请求信息"
// @Success 200 {object} updatePasswordResponse
// @Failure 400 {object} code.Failure
// @Router /admin/updatePassword [post]
func (h *handler) UpdatePassword() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
