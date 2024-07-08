package pms_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type simpleListRequest struct {
	Keyword string `form:"keyword"`
}

type simpleListResponse struct {
	List []dto.PmsProduct `json:",inline"`
}

// SimpleList 根据商品名称或货号模糊查询
// @Summary 根据商品名称或货号模糊查询
// @Description 根据商品名称或货号模糊查询
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body simpleListRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.PmsProduct}
// @Failure 400 {object} code.Failure
// @Router /product/simpleList [get]
func (h *handler) SimpleList(ctx *gin.Context) {
	req := new(simpleListRequest)
	res := new(simpleListResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.SimpleList(ctx, req.Keyword)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
