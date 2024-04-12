package ums_admin

import (
	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改指定用户信息
// @Summary 修改指定用户信息
// @Description 修改指定用户信息
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateResponse}
// @Failure 400 {object} code.Failure
// @Router /admin/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {

}
