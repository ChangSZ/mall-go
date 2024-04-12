package ums_resource

import (
	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 根据ID删除后台资源
// @Summary 根据ID删除后台资源
// @Description 根据ID删除后台资源
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /resource/delete/{id} [post]
func (h *handler) Delete(ctx *gin.Context) {

}
