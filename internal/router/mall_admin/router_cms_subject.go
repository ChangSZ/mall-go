package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/cms_subject"
	"github.com/ChangSZ/mall-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 商品专题管理
func setCmsSubjectRouter(eng *gin.Engine) {
	subjectHandler := cms_subject.New()
	subject := eng.Group("/subject", middleware.CheckToken(), middleware.DynamicAccess())
	{
		subject.GET("/listAll", subjectHandler.ListAll) // 获取全部商品专题
		subject.GET("/list", subjectHandler.List)       // 根据专题名称分页获取商品专题
	}
}
