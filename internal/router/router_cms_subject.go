package router

import (
	"github.com/ChangSZ/mall-go/internal/api/cms_subject"

	"github.com/gin-gonic/gin"
)

// 商品专题管理
func setCmsSubjectRouter(eng *gin.Engine) {
	subjectHandler := cms_subject.New()
	subject := eng.Group("/subject")
	{
		subject.GET("/listAll", subjectHandler.ListAll) // 获取全部商品专题
		subject.GET("/list", subjectHandler.List)       // 根据专题名称分页获取商品专题
	}
}
