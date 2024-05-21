package ums_menu

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateHiddenRequest struct {
	Hidden int32 `json:"hidden"`
}

type updateHiddenResponse struct {
	Count int64 `json:"count"`
}

// UpdateHidden 修改菜单显示状态
// @Summary 修改菜单显示状态
// @Description 修改菜单显示状态
// @Tags UmsMenuController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateHiddenRequest true "请求信息"
// @Success 200 {object} code.Success{data=updateHiddenResponse}
// @Failure 400 {object} code.Failure
// @Router /menu/updateHidden/{id} [post]
func (h *handler) UpdateHidden(ctx *gin.Context) {
	req := new(updateHiddenRequest)
	res := new(updateHiddenResponse)
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

	cnt, err := h.service.UpdateHidden(ctx, uri.Id, req.Hidden)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
