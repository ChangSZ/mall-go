package tool

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type hashIdsEncodeRequest struct {
	Id int64 `uri:"id"` // 需加密的数字
}

type hashIdsEncodeResponse struct {
	Val string `json:"val"` // 加密后的值
}

// HashIdsEncode HashIds 加密
// @Summary HashIds 加密
// @Description HashIds 加密
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "需加密的数字"
// @Success 200 {object} hashIdsEncodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/encode/{id} [get]
// @Security LoginToken
func (h *handler) HashIdsEncode(ctx *gin.Context) {
	req := new(hashIdsEncodeRequest)
	res := new(hashIdsEncodeResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(req.Id)})
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsEncodeError, err)
		return
	}

	res.Val = hashId

	api.ResponseOK(ctx, res)
}
