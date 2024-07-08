package minio

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type uploadRequest struct{}

type uploadResponse struct {
	Url  string `json:"url"`  // 文件访问URL
	Name string `json:"name"` // 文件名称
}

// Upload 文件上传
// @Summary 文件上传
// @Description 文件上传
// @Tags MinioController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body uploadRequest true "请求信息"
// @Success 200 {object} code.Success{data=uploadResponse}
// @Failure 400 {object} code.Failure
// @Router /minio/upload [post]
func (h *handler) Upload(ctx *gin.Context) {
	_ = new(uploadRequest)
	res := new(uploadResponse)
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err)
		return
	}
	defer file.Close()

	filename := ctx.Request.Header.Get("filename")

	url, name, err := h.service.Upload(ctx, file, filename)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Url = url
	res.Name = name
	api.Success(ctx, res)
}
