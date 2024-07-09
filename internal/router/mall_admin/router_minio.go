package mall_admin

import (
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api/mall_admin/minio"
	"github.com/ChangSZ/mall-go/internal/middleware"
)

// MinIO对象存储管理
func setMinioRouter(eng *gin.Engine) {
	handler := minio.New()
	group := eng.Group("/minio", middleware.CheckToken(), middleware.DynamicAccess())
	{
		group.POST("/upload", handler.Upload)             // 文件上传
		group.POST("/delete", handler.Delete)             // 文件删除
		group.GET("/presigned-url", handler.PresignedURL) // 文件预签名URL
	}
}
