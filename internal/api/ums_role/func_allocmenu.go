package ums_role

import (
	"github.com/gin-gonic/gin"
)

type allocMenuRequest struct{}

type allocMenuResponse struct{}

// AllocMenu 给角色分配菜单
// @Summary 给角色分配菜单
// @Description 给角色分配菜单
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body allocMenuRequest true "请求信息"
// @Success 200 {object} code.Success{data=allocMenuResponse}
// @Failure 400 {object} code.Failure
// @Router /role/allocMenu [post]
func (h *handler) AllocMenu(ctx *gin.Context) {

}
