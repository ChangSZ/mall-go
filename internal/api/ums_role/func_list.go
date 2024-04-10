package ums_role

import (
	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 根据角色名称分页获取角色列表
// @Summary 根据角色名称分页获取角色列表
// @Description 根据角色名称分页获取角色列表
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /role/list [get]
func (h *handler) List(ctx *gin.Context) {

}
