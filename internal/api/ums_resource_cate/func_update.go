package ums_resource_cate

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改后台资源分类
// @Summary 修改后台资源分类
// @Description 修改后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/update/{id} [post]
func (h *handler) Update() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
