package sms_home_recommend_product

import (
	"github.com/ChangSZ/mall-go/internal/services/sms_home_recommend_product"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加首页推荐
	// @Tags SmsHomeRecommendProductController
	// @Router /home/recommendProduct [post]
	Create(*gin.Context)

	// UpdateSort 修改推荐排序
	// @Tags SmsHomeRecommendProductController
	// @Router /home/recommendProduct/update/sort/{id} [post]
	UpdateSort(*gin.Context)

	// UpdateRecommendStatus 批量修改推荐状态
	// @Tags SmsHomeRecommendProductController
	// @Router /home/recommendProduct/update/recommendStatus [post]
	UpdateRecommendStatus(*gin.Context)

	// Delete 批量删除推荐
	// @Tags SmsHomeRecommendProductController
	// @Router /home/recommendProduct/delete [post]
	Delete(*gin.Context)

	// List 分页查询推荐
	// @Tags SmsHomeRecommendProductController
	// @Router /home/recommendProduct/list [get]
	List(*gin.Context)
}

type handler struct {
	smsHomeRecommendProductService sms_home_recommend_product.Service
}

func New() Handler {
	return &handler{
		smsHomeRecommendProductService: sms_home_recommend_product.New(),
	}
}

func (h *handler) i() {}
