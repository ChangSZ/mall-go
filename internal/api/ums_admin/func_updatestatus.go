package ums_admin

import (
	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct{}

type updateStatusResponse struct{}

// UpdateStatus 修改帐号状态
// @Summary 修改帐号状态
// @Description 修改帐号状态
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateStatusResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/updateStatus/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {

}
