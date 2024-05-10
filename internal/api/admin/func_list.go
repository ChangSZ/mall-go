package admin

import (
	"net/http"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/code"
	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/timeutil"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int    `form:"page"`      // 第几页
	PageSize int    `form:"page_size"` // 每页显示条数
	Username string `form:"username"`  // 用户名
	Nickname string `form:"nickname"`  // 昵称
	Mobile   string `form:"mobile"`    // 手机号
}

type listData struct {
	Id          int    `json:"id"`           // ID
	HashID      string `json:"hashid"`       // hashid
	Username    string `json:"username"`     // 用户名
	Nickname    string `json:"nickname"`     // 昵称
	Mobile      string `json:"mobile"`       // 手机号
	IsUsed      int    `json:"is_used"`      // 是否启用 1:是 -1:否
	IsOnline    int    `json:"is_online"`    // 是否在线 1:是 -1:否
	CreatedAt   string `json:"created_at"`   // 创建时间
	CreatedUser string `json:"created_user"` // 创建人
	UpdatedAt   string `json:"updated_at"`   // 更新时间
	UpdatedUser string `json:"updated_user"` // 更新人
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// List 管理员列表
// @Summary 管理员列表
// @Description 管理员列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true "第几页" default(1)
// @Param page_size query int true "每页显示条数" default(10)
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param mobile query string false "手机号"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
// @Security LoginToken
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.ParamBindError, validator.GetValidationError(err).Error())
		return
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	searchData := new(admin.SearchData)
	searchData.Page = page
	searchData.PageSize = pageSize
	searchData.Username = req.Username
	searchData.Nickname = req.Nickname
	searchData.Mobile = req.Mobile

	resListData, err := h.service.PageList(ctx, searchData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminListError, err)
		return
	}

	resCountData, err := h.service.PageListCount(ctx, searchData)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Response(ctx, http.StatusBadRequest, code.AdminListError, err)
		return
	}
	res.Pagination.Total = cast.ToInt(resCountData)
	res.Pagination.PerPageCount = pageSize
	res.Pagination.CurrentPage = page
	res.List = make([]listData, len(resListData))

	for k, v := range resListData {
		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
		if err != nil {
			log.WithTrace(ctx).Error(err)
			api.Response(ctx, http.StatusBadRequest, code.HashIdsEncodeError, err)
			return
		}

		isOnline := -1
		if redis.Cache().Exists(ctx, configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(v.Id)) {
			isOnline = 1
		}

		data := listData{
			Id:          cast.ToInt(v.Id),
			HashID:      hashId,
			Username:    v.Username,
			Nickname:    v.Nickname,
			Mobile:      v.Mobile,
			IsUsed:      cast.ToInt(v.IsUsed),
			IsOnline:    isOnline,
			CreatedAt:   v.CreatedAt.Format(timeutil.CSTLayout),
			CreatedUser: v.CreatedUser,
			UpdatedAt:   v.UpdatedAt.Format(timeutil.CSTLayout),
			UpdatedUser: v.UpdatedUser,
		}

		res.List[k] = data
	}

	api.ResponseOK(ctx, res)
}
