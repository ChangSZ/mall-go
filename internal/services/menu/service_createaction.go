package menu

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/menu_action"
)

type CreateMenuActionData struct {
	MenuId int32  `json:"menu_id"` // 菜单栏ID
	Method string `json:"method"`  // 请求方法
	API    string `json:"api"`     // 请求地址
}

func (s *service) CreateAction(ctx context.Context, menuActionData *CreateMenuActionData) (id int32, err error) {
	model := menu_action.NewModel()
	model.MenuId = menuActionData.MenuId
	model.Method = menuActionData.Method
	model.Api = menuActionData.API
	model.CreatedUser = core.SessionUserInfo(ctx).UserName
	model.IsDeleted = -1

	id, err = model.Create(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	return
}
