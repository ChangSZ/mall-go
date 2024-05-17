package minio

import (
	"github.com/ChangSZ/mall-go/internal/services/mall_admin/minio"

	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Upload 文件上传
	// @Tags MinioController
	// @Router /minio/upload [post]
	Upload(*gin.Context)

	// Delete 文件删除
	// @Tags MinioController
	// @Router /minio/delete [get]
	Delete(*gin.Context)
}

type handler struct {
	service minio.Service
}

func New() Handler {
	return &handler{
		service: minio.New(),
	}
}

func (h *handler) i() {}
