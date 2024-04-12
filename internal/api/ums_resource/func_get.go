package ums_resource

import (
	"github.com/gin-gonic/gin"
)

type getRequest struct{}

type getResponse struct{}

// Get 根据ID获取资源详情
// @Summary 根据ID获取资源详情
// @Description 根据ID获取资源详情
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} code.Success{data=getResponse}
// @Failure 400 {object} code.Failure
// @Router /resource/{id} [get]
func (h *handler) Get(ctx *gin.Context) {

}
