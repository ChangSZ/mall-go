package pms_brand

import (
	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type listRequest struct {
	Keyword    string `form:"keyword" binding:"omitempty"`
	ShowStatus int32  `form:"showStatus" binding:"omitempty"`
	PageSize   int    `form:"pageSize,default=5" binding:"omitempty"`
	PageNum    int    `form:"pageNum,default=1" binding:"omitempty"`
}

type listResponse struct {
	*pagehelper.ListData[dto.PmsBrand] `json:",inline"`
}

// List 根据品牌名称分页获取品牌列表
// @Summary 根据品牌名称分页获取品牌列表
// @Description 根据品牌名称分页获取品牌列表
// @Tags PmsBrandController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse}
// @Failure 400 {object} code.Failure
// @Router /brand/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetError(err).Error())
		return
	}

	list, err := h.service.List(ctx, req.Keyword, req.ShowStatus, req.PageSize, req.PageNum)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}

	res.ListData = list
	api.Success(ctx, res)
}
