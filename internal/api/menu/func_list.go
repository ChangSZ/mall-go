package menu

import (
	"net/http"

	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/services/menu"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type listData struct {
	Id     int64  `json:"id"`      // ID
	HashID string `json:"hashid"`  // hashid
	Pid    int64  `json:"pid"`     // 父类ID
	Name   string `json:"name"`    // 菜单名称
	Link   string `json:"link"`    // 链接地址
	Icon   string `json:"icon"`    // 图标
	IsUsed int32  `json:"is_used"` // 是否启用 1=启用 -1=禁用
	Sort   int32  `json:"sort"`    // 排序
}

type listResponse struct {
	List []listData `json:"list"`
}

// List 菜单列表
// @Summary 菜单列表
// @Description 菜单列表
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu [get]
// @Security LoginToken
func (h *handler) List(ctx *gin.Context) {
	res := new(listResponse)
	resListData, err := h.menuService.List(ctx, new(menu.SearchData))
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.MenuListError, err)
		return
	}

	res.List = make([]listData, len(resListData))

	for k, v := range resListData {
		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.HashIdsEncodeError, err)
			return
		}

		data := listData{
			Id:     v.Id,
			HashID: hashId,
			Pid:    v.Pid,
			Name:   v.Name,
			Link:   v.Link,
			Icon:   v.Icon,
			IsUsed: v.IsUsed,
			Sort:   v.Sort,
		}

		res.List[k] = data
	}

	api.ResponseOK(ctx, res)
}
