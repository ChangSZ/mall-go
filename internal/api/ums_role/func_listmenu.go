package ums_role

import (
	"github.com/gin-gonic/gin"
)

type listMenuRequest struct{}

type listMenuResponse struct{}

// ListMenu 获取角色相关菜单
// @Summary 获取角色相关菜单
// @Description 获取角色相关菜单
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listMenuRequest true "请求信息"
// @Success 200 {object} code.Success{data=listMenuResponse}
// @Failure 400 {object} code.Failure
// @Router /role/listMenu/{roleId} [get]
func (h *handler) ListMenu(ctx *gin.Context) {

}
