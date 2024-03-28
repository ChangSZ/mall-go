package ums_resource

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type getRequest struct{}

type getResponse struct{}

// Get 根据ID获取资源详情
// @Summary 根据ID获取资源详情
// @Description 根据ID获取资源详情
// @Tags RESOURCE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} getResponse
// @Failure 400 {object} code.Failure
// @Router /resource/{id} [get]
func (h *handler) Get() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
