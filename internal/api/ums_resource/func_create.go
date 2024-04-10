package ums_resource

import (
	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加后台资源
// @Summary 添加后台资源
// @Description 添加后台资源
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /resource/create [post]
func (h *handler) Create(ctx *gin.Context) {

}
