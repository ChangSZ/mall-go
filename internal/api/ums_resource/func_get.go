package ums_resource

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getRequest struct{}

type getResponse struct {
	UmsResource `json:",inline"`
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
	uri := new(UmsResourceUri)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	item, err := h.umsResourceService.GetItem(ctx, uri.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.UmsResource = UmsResource{
		Id:          item.Id,
		CreateTime:  item.CreateTime,
		Name:        item.Name,
		Url:         item.Url,
		Description: item.Description,
		CategoryId:  item.CategoryId,
	}
	api.Success(ctx, res)
}
