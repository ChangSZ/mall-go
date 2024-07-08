package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type treeListRequest struct{}

type treeListResponse struct {
	List []dto.UmsMenuNode `json:",inline"`
}

// TreeList 树形结构返回所有菜单列表
// @Summary 树形结构返回所有菜单列表
// @Description 树形结构返回所有菜单列表
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body treeListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.UmsMenuNode}
// @Failure 400 {object} code.Failure
// @Router /menu/treeList [get]
func (h *handler) TreeList(ctx *gin.Context) {
	_ = new(treeListRequest)
	res := new(treeListResponse)
	list, err := h.service.TreeList(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
