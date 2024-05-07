package sms_home_recommend_subject

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct{}

type deleteResponse struct{}

// Delete 批量删除推荐专题
// @Summary 批量删除推荐专题
// @Description 批量删除推荐专题
// @Tags SmsHomeRecommendSubjectController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /home/recommendSubject/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	api.Success(ctx, nil)
}
