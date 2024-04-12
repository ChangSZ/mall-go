package ums_role

import (
	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct{}

// ListAll 获取所有角色
// @Summary 获取所有角色
// @Description 获取所有角色
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} code.Success{data=listAllResponse}
// @Failure 400 {object} code.Failure
// @Router /role/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {

}
