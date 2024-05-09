package router

import (
	"github.com/ChangSZ/mall-go/internal/api/oms_company_address"

	"github.com/gin-gonic/gin"
)

// 订单管理
func setOmsCompanyAddressRouter(eng *gin.Engine) {
	handler := oms_company_address.New()
	group := eng.Group("/companyAddress")
	{
		group.GET("/list", handler.List) // 获取所有收货地址
	}
}
