package router

import (
	"github.com/ChangSZ/mall-go/internal/api/pms_brand"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 商品分类管理
func setPmsBrandRouter(eng *gin.Engine) {
	brandHandler := pms_brand.New()
	brands := eng.Group("/brand", middleware.CheckToken())
	{
		brands.POST("/create", brandHandler.Create)                            // 添加品牌
		brands.POST("/update/:id", brandHandler.Update)                        // 更新品牌
		brands.POST("/delete/:id", brandHandler.Delete)                        // 删除品牌
		brands.POST("/delete/batch", brandHandler.DeleteBatch)                 // 批量删除品牌
		brands.GET("/list", brandHandler.List)                                 // 根据品牌名称分页获取品牌列表
		brands.GET("/listAll", brandHandler.ListAll)                           // 获取全部品牌列表
		brands.GET("/:id", brandHandler.GetItem)                               // 根据编号查询品牌信息
		brands.POST("/update/showStatus", brandHandler.UpdateShowStatus)       // 批量更新显示状态
		brands.POST("/update/factoryStatus", brandHandler.UpdateFactoryStatus) // 批量更新厂家制造商状态
	}
}
