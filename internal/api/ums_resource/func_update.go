package ums_resource

import (
	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改后台资源
// @Summary 修改后台资源
// @Description 修改后台资源
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /resource/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {

}
