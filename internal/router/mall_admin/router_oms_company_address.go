package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/oms_company_address"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 订单管理
func setOmsCompanyAddressRouter(eng *gin.Engine) {
	handler := oms_company_address.New()
	group := eng.Group("/companyAddress", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.GET("/list", handler.List) // 获取所有收货地址
	}
}
