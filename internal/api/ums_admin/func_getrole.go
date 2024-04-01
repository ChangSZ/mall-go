package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type getRoleRequest struct{}

type getRoleResponse struct{}

// GetRole 获取指定用户的角色
// @Summary 获取指定用户的角色
// @Description 获取指定用户的角色
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRoleRequest true "请求信息"
// @Success 200 {object} getRoleResponse
// @Failure 400 {object} code.Failure
// @Router /admin/role/{adminId} [get]
func (h *handler) GetRole() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
