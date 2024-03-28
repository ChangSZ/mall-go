package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type listResourceRequest struct{}

type listResourceResponse struct{}

// ListResource 获取角色相关资源
// @Summary 获取角色相关资源
// @Description 获取角色相关资源
// @Tags ROLE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listResourceRequest true "请求信息"
// @Success 200 {object} listResourceResponse
// @Failure 400 {object} code.Failure
// @Router /role/listResource/{roleId} [get]
func (h *handler) ListResource() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
