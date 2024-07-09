package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/cms_prefrence_area"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// 商品优选管理
func setCmsPrefrenceAreaRouter(eng *gin.Engine) {
	areaHandler := cms_prefrence_area.New()
	area := eng.Group("/prefrenceArea", middleware.CheckToken(), middleware.DynamicAccess())
	{
		area.GET("/listAll", areaHandler.ListAll) // 获取所有商品优选
	}
}
