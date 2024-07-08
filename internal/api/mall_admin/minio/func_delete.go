package minio

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	ObjectName string `form:"objectName"`
}

type deleteResponse struct{}

// Delete 文件删除
// @Summary 文件删除
// @Description 文件删除
// @Tags MinioController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} code.Success{data=deleteResponse}
// @Failure 400 {object} code.Failure
// @Router /minio/delete [post]
func (h *handler) Delete(ctx *gin.Context) {
	req := new(deleteRequest)
	_ = new(deleteResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	err := h.service.Delete(ctx, req.ObjectName)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx)
}
