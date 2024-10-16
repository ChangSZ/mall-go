package ums_resource

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type getRequest struct{}

type getResponse struct {
	dto.UmsResource `json:",inline"`
}

// Get 根据ID获取资源详情
// @Summary 根据ID获取资源详情
// @Description 根据ID获取资源详情
// @Tags UmsResourceController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getRequest true "请求信息"
// @Success 200 {object} code.Success{data=getResponse}
// @Failure 400 {object} code.Failure
// @Router /resource/{id} [get]
func (h *handler) Get(ctx *gin.Context) {
	_ = new(getRequest)
	res := new(getResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	item, err := h.service.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.UmsResource = *item
	api.Success(ctx, res)
}
