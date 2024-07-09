package ums_role

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type allocResourceRequest struct {
	RoleId      int64   `form:"role_id" binding:"required"`
	ResourceIds []int64 `form:"resourceIds"`
}

type allocResourceResponse struct {
	Count int64 `json:",inline"`
}

// AllocResource 给角色分配资源
// @Summary 给角色分配资源
// @Description 给角色分配资源
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body allocResourceRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /role/allocResource [post]
func (h *handler) AllocResource(ctx *gin.Context) {
	req := new(allocResourceRequest)
	res := new(allocResourceResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.AllocResource(ctx, req.RoleId, req.ResourceIds)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
