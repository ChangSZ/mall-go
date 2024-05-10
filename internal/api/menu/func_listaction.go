package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type listActionRequest struct {
	Id string `form:"id"` // hashID
}

type listActionData struct {
	HashId string `json:"hash_id"` // hashID
	MenuId int64  `json:"menu_id"` // 菜单栏ID
	Method string `json:"method"`  // 调用方secret
	API    string `json:"api"`     // 调用方对接人
}

type listActionResponse struct {
	MenuName string           `json:"menu_name"`
	List     []listActionData `json:"list"`
}

// ListAction 功能权限列表
// @Summary 功能权限列表
// @Description 功能权限列表
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id query string true "hashID"
// @Success 200 {object} listActionResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu_action [get]
// @Security LoginToken
func (h *handler) ListAction(ctx *gin.Context) {
	req := new(listActionRequest)
	res := new(listActionResponse)
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

	searchOneData := new(menu.SearchOneData)
	searchOneData.Id = id

	menuInfo, err := h.service.Detail(ctx, searchOneData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuDetailError, err)
		return
	}

	res.MenuName = menuInfo.Name

	searchListData := new(menu.SearchListActionData)
	searchListData.MenuId = menuInfo.Id

	resListData, err := h.service.ListAction(ctx, searchListData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AuthorizedListAPIError, err)
		return
	}

	res.List = make([]listActionData, len(resListData))

	for k, v := range resListData {
		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.HashIdsEncodeError, err)
			return
		}

		data := listActionData{
			HashId: hashId,
			MenuId: v.MenuId,
			Method: v.Method,
			API:    v.Api,
		}

		res.List[k] = data
	}

	api.ResponseOK(ctx, res)
}
