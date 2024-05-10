package sms_home_recommend_subject

import (
	"github.com/ChangSZ/mall-go/internal/services/sms_home_recommend_subject"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加首页推荐专题
	// @Tags SmsHomeRecommendSubjectController
	// @Router /home/recommendSubject [post]
	Create(*gin.Context)

	// UpdateSort 修改推荐专题排序
	// @Tags SmsHomeRecommendSubjectController
	// @Router /home/recommendSubject/update/sort/{id} [post]
	UpdateSort(*gin.Context)

	// UpdateRecommendStatus 批量修改推荐专题状态
	// @Tags SmsHomeRecommendSubjectController
	// @Router /home/recommendSubject/update/recommendStatus [post]
	UpdateRecommendStatus(*gin.Context)

	// Delete 批量删除推荐专题
	// @Tags SmsHomeRecommendSubjectController
	// @Router /home/recommendSubject/delete [post]
	Delete(*gin.Context)

	// List 分页查询推荐专题
	// @Tags SmsHomeRecommendSubjectController
	// @Router /home/recommendSubject/list [get]
	List(*gin.Context)
}

type handler struct {
	service sms_home_recommend_subject.Service
}

func New() Handler {
	return &handler{
		service: sms_home_recommend_subject.New(),
	}
}

func (h *handler) i() {}
