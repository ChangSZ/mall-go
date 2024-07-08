package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type listRequest struct{}

type listResponse struct {
	List []dto.SmsFlashPromotionSession `json:",inline"`
}

// List 获取全部场次
// @Summary 获取全部场次
// @Description 获取全部场次
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.SmsFlashPromotionSession}
// @Failure 400 {object} code.Failure
// @Router /flashSession/list [get]
func (h *handler) List(ctx *gin.Context) {
	_ = new(listRequest)
	res := new(listResponse)

	list, err := h.service.ListAll(ctx)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
