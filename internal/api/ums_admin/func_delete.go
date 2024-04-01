package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 删除指定用户信息
// @Summary 删除指定用户信息
// @Description 删除指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /admin/delete/{id} [post]
func (h *handler) Delete() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
