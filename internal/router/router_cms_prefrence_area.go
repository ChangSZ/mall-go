package router

import (
	"github.com/ChangSZ/mall-go/internal/api/cms_prefrence_area"

	"github.com/gin-gonic/gin"
)

// 商品优选管理
func setCmsPrefrenceAreaRouter(eng *gin.Engine) {
	areaHandler := cms_prefrence_area.New()
	area := eng.Group("/prefrenceArea")
	{
		area.GET("/listAll", areaHandler.ListAll) // 获取所有商品优选
	}
}
