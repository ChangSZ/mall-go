package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updateRoleRequest struct{}

type updateRoleResponse struct{}

// UpdateRole 给用户分配角色
// @Summary 给用户分配角色
// @Description 给用户分配角色
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRoleRequest true "请求信息"
// @Success 200 {object} updateRoleResponse
// @Failure 400 {object} code.Failure
// @Router /admin/role/update [post]
func (h *handler) UpdateRole() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
