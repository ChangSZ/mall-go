package ums_menu

import (
	"github.com/gin-gonic/gin"
)

type updateRequest struct{}

type updateResponse struct{}

// Update 修改后台菜单
// @Summary 修改后台菜单
// @Description 修改后台菜单
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} updateResponse
// @Failure 400 {object} code.Failure
// @Router /menu/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {

}
