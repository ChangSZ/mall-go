package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改指定用户信息
// @Summary 修改指定用户信息
// @Description 修改指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /admin/update/{id} [post]
func (h *handler) Update() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
