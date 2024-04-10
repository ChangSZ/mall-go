package ums_menu

import (
	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct{}

// List 分页查询后台菜单
// @Summary 分页查询后台菜单
// @Description 分页查询后台菜单
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /menu/list/{parentId} [get]
func (h *handler) List(ctx *gin.Context) {

}
