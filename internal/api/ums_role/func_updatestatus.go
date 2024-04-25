package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct {
	Status int32 `json:"status"`
}

type updateStatusResponse struct {
	Count int64 `json:",inline"`
}

// UpdateStatus 修改角色状态
// @Summary 修改角色状态
// @Description 修改角色状态
// @Tags UmsRoleController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /role/updateStatus/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	req := new(updateStatusRequest)
	res := new(updateStatusResponse)
	uri := new(UmsRoleUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data := &ums_role.UmsRole{
		Status: req.Status,
	}
	cnt, err := h.umsRoleService.Update(ctx, uri.Id, data)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新status个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
