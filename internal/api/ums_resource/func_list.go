package ums_resource

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type listRequest struct{}

type listResponse struct{}

// List 分页模糊查询后台资源
// @Summary 分页模糊查询后台资源
// @Description 分页模糊查询后台资源
// @Tags RESOURCE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /resource/list [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
