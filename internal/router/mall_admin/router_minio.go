package mall_admin

import (
	"github.com/ChangSZ/mall-go/internal/api/mall_admin/minio"

	"github.com/gin-gonic/gin"
)

// MinIO对象存储管理
func setMinioRouter(eng *gin.Engine) {
	handler := minio.New()
	group := eng.Group("/minio")
	{
		group.POST("/upload", handler.Upload) // 文件上传
		group.POST("/delete", handler.Delete) // 文件删除
	}
}
