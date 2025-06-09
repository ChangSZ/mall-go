package sms_coupon_history

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	CouponId  int64  `form:"couponId" binding:"omitempty"`
	UseStatus int32  `form:"useStatus" binding:"omitempty"`
	OrderSn   string `form:"orderSn" binding:"omitempty"`
	PageSize  int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum   int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int                    `json:"pageNum"`
	PageSize  int                    `json:"pageSize"`
	TotalPage int64                  `json:"totalPage"`
	Total     int64                  `json:"total"`
	List      []dto.SmsCouponHistory `json:"list"`
}

// List 根据优惠券id，使用状态，订单编号分页获取领取记录
// @Summary 根据优惠券id，使用状态，订单编号分页获取领取记录
// @Description 根据优惠券id，使用状态，订单编号分页获取领取记录
// @Tags SmsCouponHistoryController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /couponHistory/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, total, err := h.service.List(
		ctx, req.CouponId, req.UseStatus, req.OrderSn, req.PageSize, req.PageNum)
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
