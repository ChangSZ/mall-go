package sms_home_brand

import (
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加首页推荐品牌
	// @Tags SmsHomeBrandController
	// @Router /home/brand [post]
	Create(*gin.Context)

	// UpdateSort 修改推荐品牌排序
	// @Tags SmsHomeBrandController
	// @Router /home/brand/update/sort/{id} [post]
	UpdateSort(*gin.Context)

	// UpdateRecommendStatus 批量修改推荐品牌状态
	// @Tags SmsHomeBrandController
	// @Router /home/brand/update/recommendStatus [post]
	UpdateRecommendStatus(*gin.Context)

	// Delete 批量删除推荐品牌
	// @Tags SmsHomeBrandController
	// @Router /home/brand/delete [post]
	Delete(*gin.Context)

	// List 分页查询推荐品牌
	// @Tags SmsHomeBrandController
	// @Router /home/brand/list [get]
	List(*gin.Context)
}

type handler struct {
	// smsHomeBrandService sms_home_brand.Service
}

func New() Handler {
	return &handler{
		// smsHomeBrandService: sms_home_brand.New(),
	}
}

func (h *handler) i() {}
