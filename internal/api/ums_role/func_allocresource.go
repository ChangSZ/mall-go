package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type allocResourceRequest struct{}

type allocResourceResponse struct{}

// AllocResource 给角色分配资源
// @Summary 给角色分配资源
// @Description 给角色分配资源
// @Tags ROLE
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body allocResourceRequest true "请求信息"
// @Success 200 {object} allocResourceResponse
// @Failure 400 {object} code.Failure
// @Router /role/allocResource [post]
func (h *handler) AllocResource() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
