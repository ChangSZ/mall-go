package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加角色
// @Summary 添加角色
// @Description 添加角色
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /role/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
