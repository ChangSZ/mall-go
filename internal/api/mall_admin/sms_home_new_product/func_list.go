package sms_home_new_product

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	ProductName     string `form:"productName" binding:"omitempty"`
	RecommendStatus int32  `form:"recommendStatus" binding:"omitempty"`
	PageSize        int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum         int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	*pagehelper.ListData[dto.SmsHomeNewProduct] `json:",inline"`
}

// List 分页查询首页新品
// @Summary 分页查询首页新品
// @Description 分页查询首页新品
// @Tags SmsHomeNewProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /home/newProduct/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.List(
		ctx, req.ProductName, req.RecommendStatus, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ListData = list
	api.Success(ctx, res)
}
