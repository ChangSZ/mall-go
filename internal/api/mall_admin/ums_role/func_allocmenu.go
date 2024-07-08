package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type allocMenuRequest struct {
	RoleId  int64  `form:"roleId" binding:"required"`
	MenuIds string `form:"menuIds"`
}

type allocMenuResponse struct {
	Count int64 `json:",inline"`
}

// AllocMenu 给角色分配菜单
// @Summary 给角色分配菜单
// @Description 给角色分配菜单
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body allocMenuRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /role/allocMenu [post]
func (h *handler) AllocMenu(ctx *gin.Context) {
	req := new(allocMenuRequest)
	res := new(allocMenuResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.AllocMenu(ctx, req.RoleId, req.MenuIds)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
