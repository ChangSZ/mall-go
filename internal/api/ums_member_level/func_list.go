package ums_member_level

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type listRequest struct{}

type listResponse struct{}

// List 查询所有会员等级
// @Summary 查询所有会员等级
// @Description 查询所有会员等级
// @Tags UmsMemberLevelController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /memberLevel/list [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
