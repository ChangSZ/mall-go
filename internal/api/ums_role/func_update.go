package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改角色
// @Summary 修改角色
// @Description 修改角色
// @Tags ROLE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /role/update/{id} [post]
func (h *handler) Update() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
