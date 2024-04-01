package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type treeListRequest struct{}

type treeListResponse struct{}

// TreeList 树形结构返回所有菜单列表
// @Summary 树形结构返回所有菜单列表
// @Description 树形结构返回所有菜单列表
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body treeListRequest true "请求信息"
// @Success 200 {object} treeListResponse
// @Failure 400 {object} code.Failure
// @Router /menu/treeList [get]
func (h *handler) TreeList() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}