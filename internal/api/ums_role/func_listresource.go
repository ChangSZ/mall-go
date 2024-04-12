package ums_role

import (
	"github.com/gin-gonic/gin"
)

type listResourceRequest struct{}

type listResourceResponse struct{}

// ListResource 获取角色相关资源
// @Summary 获取角色相关资源
// @Description 获取角色相关资源
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listResourceRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResourceResponse}
// @Failure 400 {object} code.Failure
// @Router /role/listResource/{roleId} [get]
func (h *handler) ListResource(ctx *gin.Context) {

}
