package authorized

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
)

type createAPIRequest struct {
	Id     string `form:"id"`     // HashID
	Method string `form:"method"` // 请求方法
	API    string `form:"api"`    // 请求地址
}

type createAPIResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// CreateAPI 授权调用方接口地址
// @Summary 授权调用方接口地址
// @Description 授权调用方接口地址
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "HashID"
// @Param method formData string true "请求方法"
// @Param api formData string true "请求地址"
// @Success 200 {object} createAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [post]
// @Security LoginToken
func (h *handler) CreateAPI(ctx *gin.Context) {
	req := new(createAPIRequest)
	res := new(createAPIResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	ids, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	id := int64(ids[0])
	// 通过 id 查询出 business_key
	authorizedInfo, err := h.service.Detail(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedDetailError, err)
		return
	}

	createAPIData := new(authorized.CreateAuthorizedAPIData)
	createAPIData.BusinessKey = authorizedInfo.BusinessKey
	createAPIData.Method = req.Method
	createAPIData.API = req.API

	createId, err := h.service.CreateAPI(ctx, createAPIData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedCreateAPIError, err)
		return
	}

	res.Id = createId
	api.ResponseOK(ctx, res)
}
