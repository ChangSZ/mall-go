package ums_role

import (
	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 修改角色状态
// @Summary 修改角色状态
// @Description 修改角色状态
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /role/updateStatus/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {

}
