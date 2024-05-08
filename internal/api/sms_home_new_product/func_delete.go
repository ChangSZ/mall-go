package sms_home_new_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	Ids []int64 `form:"ids" binding:"required"`
}

type deleteResponse struct {
	Count int64 `json:",inline"`
}

// Delete 批量删除首页新品
// @Summary 批量删除首页新品
// @Description 批量删除首页新品
// @Tags SmsHomeNewProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /home/newProduct/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	res := new(deleteResponse)
	cnt, err := h.smsHomeNewProductService.Delete(ctx, req.Ids)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	if cnt == 0 {
		api.Failed(ctx, "删除个数为0")
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
