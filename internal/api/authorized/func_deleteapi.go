package authorized

import (
	"net/http"

	"github.com/ChangSZ/golib/log"
	"github.com/gin-gonic/gin"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/pkg/validator"
)

type deleteAPIRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteAPIResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// DeleteAPI 删除调用方接口地址
// @Summary 删除调用方接口地址
// @Description 删除调用方接口地址
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param id path string true "主键ID"
// @Success 200 {object} deleteAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api/{id} [delete]
// @Security LoginToken
func (h *handler) DeleteAPI(ctx *gin.Context) {
	req := new(deleteAPIRequest)
	res := new(deleteAPIResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
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
	err = h.service.DeleteAPI(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedDeleteAPIError, err)
		return
	}

	res.Id = id
	api.ResponseOK(ctx, res)
}
