package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type listAllRequest struct{}

type listAllResponse struct{}

// ListAll 获取所有角色
// @Summary 获取所有角色
// @Description 获取所有角色
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} listAllResponse
// @Failure 400 {object} code.Failure
// @Router /admin/listAll [get]
func (h *handler) ListAll() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
