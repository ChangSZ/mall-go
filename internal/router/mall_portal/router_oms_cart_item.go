package mall_portal

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_portal/oms_cart_item"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 购物车管理
func setOmsCartItemBrandRouter(eng *gin.Engine) {
	handler := oms_cart_item.New()
	group := eng.Group("/cart", middleware.CheckMemberToken())
	{
		group.POST("/add", handler.Add)                             // 添加商品到购物车
		group.GET("/list", handler.List)                            // 获取当前会员的购物车列表
		group.GET("/list/promotion", handler.ListPromotion)         // 获取当前会员的购物车列表,包括促销信息
		group.GET("/update/quantity", handler.UpdateQuantity)       // 修改购物车中指定商品的数量（我觉着应该是POST）
		group.GET("/getProduct/:productId", handler.GetCartProduct) // 获取购物车中指定商品的规格,用于重选规格
		group.POST("/update/attr", handler.UpdateAttr)              // 修改购物车中商品的规格
		group.POST("/delete", handler.Delete)                       // 删除购物车中的指定商品
		group.POST("/clear", handler.Clear)                         // 清空当前会员的购物车
	}
}
