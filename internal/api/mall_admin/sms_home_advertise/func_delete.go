package sms_home_advertise

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

// Delete 删除广告
// @Summary 删除广告
// @Description 删除广告
// @Tags SmsHomeAdvertiseController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /home/advertise/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	res := new(deleteResponse)
	cnt, err := h.service.Delete(ctx, req.Ids)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
