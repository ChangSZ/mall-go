package sms_home_recommend_subject

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

// Delete 批量删除推荐专题
// @Summary 批量删除推荐专题
// @Description 批量删除推荐专题
// @Tags SmsHomeRecommendSubjectController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /home/recommendSubject/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	res := new(deleteResponse)
	cnt, err := h.smsHomeRecommendSubjectService.Delete(ctx, req.Ids)
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
