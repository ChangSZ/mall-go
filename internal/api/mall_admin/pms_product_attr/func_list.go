package pms_product_attr

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	AttrType int32 `json:"type" binding:"omitempty"`
	PageSize int   `form:"pageSize,default=5" binding:"omitempty"`
	PageNum  int   `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	*pagehelper.ListData[dto.PmsProductAttribute] `json:",inline"`
}

// List 根据分类查询属性列表或参数列表
// @Summary 根据分类查询属性列表或参数列表
// @Description 根据分类查询属性列表或参数列表
// @Tags PmsProductAttributeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /productAttribute/list/{cid} [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	uri := new(dto.PmsProductAttrUri)
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

	list, err := h.service.List(ctx, uri.Cid, req.AttrType, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}
