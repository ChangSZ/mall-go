package mall_portal

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_portal/ums_member_receive_address"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 会员收货地址管理
func setUmsMemberReceiveAddressRouter(eng *gin.Engine) {
	handler := ums_member_receive_address.New()
	group := eng.Group("/member/address", middleware.CheckMemberToken())
	{
		group.POST("/add", handler.Add)           // 添加收货地址
		group.POST("/delete/:id", handler.Delete) // 删除收货地址
		group.POST("/update/:id", handler.Update) // 修改收货地址
		group.GET("/list", handler.List)          // 获取所有收货地址
		group.GET("/:id", handler.GetItem)        // 获取收货地址详情
	}
}
