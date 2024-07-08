package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type updateStatusRequest struct {
	Status int32 `json:"status"`
}

type updateStatusResponse struct {
	Count int64 `json:"count"`
}

// UpdateStatus 修改帐号状态
// @Summary 修改帐号状态
// @Description 修改帐号状态
// @Tags UmsAdminController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateStatusRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /admin/updateStatus/{id} [post]
func (h *handler) UpdateStatus(ctx *gin.Context) {
	req := new(updateStatusRequest)
	res := new(updateStatusResponse)
	uri := new(dto.UriID)
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

	cnt, err := h.service.UpdateStatus(ctx, uri.Id, req.Status)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "更新状态失败, 未找到相关数据或状态无需变更")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
