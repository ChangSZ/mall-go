package ums_role

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
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
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	cnt, err := h.service.UpdateStatus(ctx, uri.Id, req.Status)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
