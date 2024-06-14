package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	// 订单状态：-1->全部；0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭
	Status   int32 `form:"status,default=-1" binding:"omitempty,oneof=-1 0 1 2 3 4"`
	PageNum  int   `form:"pageNum,default=1" binding:"omitempty"`
	PageSize int   `form:"pageSize,default=5" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int               `json:"pageNum"`
	PageSize  int               `json:"pageSize"`
	TotalPage int64             `json:"totalPage"`
	Total     int64             `json:"total"`
	List      []dto.OrderDetail `json:"list"`
}

// List 按状态分页获取用户订单列表
// @Summary 按状态分页获取用户订单列表
// @Description 按状态分页获取用户订单列表
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /order/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, total, err := h.service.List(ctx, req.Status, req.PageNum, req.PageSize)
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
