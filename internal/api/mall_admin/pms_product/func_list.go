package pms_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	dto.PmsProductQueryParam
	PageSize int `form:"pageSize,default=5" binding:"omitempty"`
	PageNum  int `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int              `json:"pageNum"`
	PageSize  int              `json:"pageSize"`
	TotalPage int64            `json:"totalPage"`
	Total     int64            `json:"total"`
	List      []dto.PmsProduct `json:"list"`
}

// List 查询商品
// @Summary 查询商品
// @Description 查询商品
// @Tags PmsProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /product/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, total, err := h.service.List(ctx, req.PmsProductQueryParam, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize
	totalPage := total / int64(req.PageSize)
	if total%int64(req.PageSize) > 0 {
		totalPage += 1
	}
	res.TotalPage = totalPage
	res.Total = total
	res.List = list
	api.Success(ctx, res)
}
