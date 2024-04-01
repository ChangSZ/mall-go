package ums_resource_cate

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type listAllRequest struct{}

type listAllResponse struct{}

// ListAll 查询所有后台资源分类
// @Summary 查询所有后台资源分类
// @Description 查询所有后台资源分类
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} listAllResponse
// @Failure 400 {object} code.Failure
// @Router /resourceCategory/listAll [get]
func (h *handler) ListAll() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
