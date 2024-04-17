package authorized

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type listAPIRequest struct {
	Id string `form:"id"` // hashID
}

type listAPIData struct {
	HashId      string `json:"hash_id"`      // hashID
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 调用方secret
	API         string `json:"api"`          // 调用方对接人
}

type listAPIResponse struct {
	BusinessKey string        `json:"business_key"` // 调用方key
	List        []listAPIData `json:"list"`
}

// ListAPI 调用方接口地址列表
// @Summary 调用方接口地址列表
// @Description 调用方接口地址列表
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id query string true "hashID"
// @Success 200 {object} listAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [get]
// @Security LoginToken
func (h *handler) ListAPI(ctx *gin.Context) {
	req := new(listAPIRequest)
	res := new(listAPIResponse)
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

	id := int32(ids[0])

	// 通过 id 查询出 business_key
	authorizedInfo, err := h.authorizedService.Detail(ctx, id)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedDetailError, err)
		return
	}

	res.BusinessKey = authorizedInfo.BusinessKey

	searchAPIData := new(authorized.SearchAPIData)
	searchAPIData.BusinessKey = authorizedInfo.BusinessKey

	resListData, err := h.authorizedService.ListAPI(ctx, searchAPIData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedListAPIError, err)
		return
	}

	res.List = make([]listAPIData, len(resListData))

	for k, v := range resListData {
		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.HashIdsEncodeError, err)
			return
		}

		data := listAPIData{
			HashId:      hashId,
			BusinessKey: v.BusinessKey,
			Method:      v.Method,
			API:         v.Api,
		}

		res.List[k] = data
	}

	api.ResponseOK(ctx, res)
}
