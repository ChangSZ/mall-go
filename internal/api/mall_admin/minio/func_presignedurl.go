package minio

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type presignedURLRequest struct {
	Bucket     string `form:"bucket"`
	ObjectName string `form:"objectName"`
}

type presignedURLResponse struct{}

// PresignedURL 文件预签名URL
// @Summary 文件预签名URL
// @Description 文件预签名URL
// @Tags MinioController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body presignedURLRequest true "请求信息"
// @Success 200 {object} code.Success{data=string}
// @Failure 400 {object} code.Failure
// @Router /minio/presigned-url [get]
func (h *handler) PresignedURL(ctx *gin.Context) {
	req := new(presignedURLRequest)
	_ = new(presignedURLResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	if req.Bucket == "" {
		api.Failed(ctx, "非预期的URL")
		return
	}

	url, err := h.service.PresignedURL(ctx, req.Bucket, req.ObjectName)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, url)
}
