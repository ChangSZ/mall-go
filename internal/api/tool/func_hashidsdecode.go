package tool

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type hashIdsDecodeRequest struct {
	Id string `uri:"id"` // 需解密的密文
}

type hashIdsDecodeResponse struct {
	Val int `json:"val"` // 解密后的值
}

// HashIdsDecode HashIds 解密
// @Summary HashIds 解密
// @Description HashIds 解密
// @Tags API.tool
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "需解密的密文"
// @Success 200 {object} hashIdsDecodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/hashids/decode/{id} [get]
// @Security LoginToken
func (h *handler) HashIdsDecode(ctx *gin.Context) {
	req := new(hashIdsDecodeRequest)
	res := new(hashIdsDecodeResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	hashId, err := h.hashids.HashidsDecode(req.Id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.HashIdsDecodeError, err)
		return
	}

	res.Val = hashId[0]

	api.ResponseOK(ctx, res)
}
