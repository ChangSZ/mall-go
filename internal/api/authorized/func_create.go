package authorized

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	BusinessKey       string `form:"business_key"`       // 调用方key
	BusinessDeveloper string `form:"business_developer"` // 调用方对接人
	Remark            string `form:"remark"`             // 备注
}

type createResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Create 新增调用方
// @Summary 新增调用方
// @Description 新增调用方
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param business_key formData string true "调用方key"
// @Param business_developer formData string true "调用方对接人"
// @Param remark formData string true "备注"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized [post]
// @Security LoginToken
func (h *handler) Create(ctx *gin.Context) {
	req := new(createRequest)
	res := new(createResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	createData := new(authorized.CreateAuthorizedData)
	createData.BusinessKey = req.BusinessKey
	createData.BusinessDeveloper = req.BusinessDeveloper
	createData.Remark = req.Remark

	id, err := h.service.Create(ctx, createData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedCreateError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}
