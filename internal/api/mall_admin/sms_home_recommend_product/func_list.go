package sms_home_recommend_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	ProductName     string `form:"productName" binding:"omitempty"`
	RecommendStatus int32  `form:"recommendStatus" binding:"omitempty"`
	PageSize        int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum         int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	PageNum   int                           `json:"pageNum"`
	PageSize  int                           `json:"pageSize"`
	TotalPage int64                         `json:"totalPage"`
	Total     int64                         `json:"total"`
	List      []dto.SmsHomeRecommendProduct `json:"list"`
}

// List 分页查询推荐
// @Summary 分页查询推荐
// @Description 分页查询推荐
// @Tags SmsHomeRecommendProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /home/recommendProduct/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, total, err := h.service.List(
		ctx, req.ProductName, req.RecommendStatus, req.PageSize, req.PageNum)
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
