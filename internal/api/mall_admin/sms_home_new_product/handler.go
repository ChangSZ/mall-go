package sms_home_new_product

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/sms_home_new_product"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 添加首页新品
	// @Tags SmsHomeNewProductController
	// @Router /home/newProduct [post]
	Create(*gin.Context)

	// UpdateSort 修改首页新品排序
	// @Tags SmsHomeNewProductController
	// @Router /home/newProduct/update/sort/{id} [post]
	UpdateSort(*gin.Context)

	// UpdateRecommendStatus 批量修改首页新品状态
	// @Tags SmsHomeNewProductController
	// @Router /home/newProduct/update/recommendStatus [post]
	UpdateRecommendStatus(*gin.Context)

	// Delete 批量删除首页新品
	// @Tags SmsHomeNewProductController
	// @Router /home/newProduct/delete [post]
	Delete(*gin.Context)

	// List 分页查询首页新品
	// @Tags SmsHomeNewProductController
	// @Router /home/newProduct/list [get]
	List(*gin.Context)
}

type handler struct {
	service sms_home_new_product.Service
}

func New() Handler {
	return &handler{
		service: sms_home_new_product.New(),
	}
}

func (h *handler) i() {}
