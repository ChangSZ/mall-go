package ums_admin

import (
	"github.com/gin-gonic/gin"
)

type updateRoleRequest struct{}

type updateRoleResponse struct{}

// UpdateRole 给用户分配角色
// @Summary 给用户分配角色
// @Description 给用户分配角色
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRoleRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateRoleResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/role/update [post]
func (h *handler) UpdateRole(ctx *gin.Context) {

}
