package ums_role

import (
	"github.com/gin-gonic/gin"
)

type allocResourceRequest struct{}

type allocResourceResponse struct{}

// AllocResource 给角色分配资源
// @Summary 给角色分配资源
// @Description 给角色分配资源
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body allocResourceRequest true "请求信息"
// @Success 200 {object} code.Success{data=allocResourceResponse}
// @Failure 400 {object} code.Failure
// @Router /role/allocResource [post]
func (h *handler) AllocResource(ctx *gin.Context) {

}
