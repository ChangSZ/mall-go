package ums_resource

import (
	"github.com/gin-gonic/gin"
)

type listAllRequest struct{}

type listAllResponse struct{}

// ListAll 查询所有后台资源
// @Summary 查询所有后台资源
// @Description 查询所有后台资源
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listAllRequest true "请求信息"
// @Success 200 {object} listAllResponse
// @Failure 400 {object} code.Failure
// @Router /resource/listAll [get]
func (h *handler) ListAll(ctx *gin.Context) {

}
