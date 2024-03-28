package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 修改角色状态
// @Summary 修改角色状态
// @Description 修改角色状态
// @Tags ROLE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} updateStatusResponse
// @Failure 400 {object} code.Failure
// @Router /role/updateStatus/{id} [post]
func (h *handler) UpdateStatus() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
