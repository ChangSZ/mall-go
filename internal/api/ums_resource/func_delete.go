package ums_resource

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 根据ID删除后台资源
// @Summary 根据ID删除后台资源
// @Description 根据ID删除后台资源
// @Tags RESOURCE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /resource/delete/{id} [post]
func (h *handler) Delete() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
