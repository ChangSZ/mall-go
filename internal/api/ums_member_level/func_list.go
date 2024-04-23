package ums_member_level

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	DefaultStatus int32 `form:"defaultStatus"`
}

type listResponse struct {
	List []UmsMemberLevel `json:",inline"`
}

// List 查询所有会员等级
// @Summary 查询所有会员等级
// @Description 查询所有会员等级
// @Tags UmsMemberLevelController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=listResponse.List}
// @Failure 400 {object} code.Failure
// @Router /memberLevel/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	res := new(listResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	list, err := h.umsMemberLevelService.List(ctx, req.DefaultStatus)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = make([]UmsMemberLevel, 0, len(list))
	for _, v := range list {
		res.List = append(res.List, UmsMemberLevel{
			Id:                    v.Id,
			Name:                  v.Name,
			GrowthPoint:           v.GrowthPoint,
			DefaultStatus:         v.DefaultStatus,
			FreeFreightPoint:      v.FreeFreightPoint,
			CommentGrowthPoint:    v.CommentGrowthPoint,
			PriviledgeFreeFreight: v.PriviledgeFreeFreight,
			PriviledgeSignIn:      v.PriviledgeSignIn,
			PriviledgeComment:     v.PriviledgeComment,
			PriviledgePromotion:   v.PriviledgePromotion,
			PriviledgeMemberPrice: v.PriviledgeMemberPrice,
			PriviledgeBirthday:    v.PriviledgeBirthday,
			Note:                  v.Note,
		})
	}
	api.Success(ctx, res.List)
}
