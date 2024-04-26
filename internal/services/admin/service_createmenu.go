package admin

import (
	"context"
	"strings"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/admin_menu"

	"github.com/spf13/cast"
)

type CreateMenuData struct {
	AdminId int64  `form:"admin_id"` // AdminID
	Actions string `form:"actions"`  // 功能权限ID,多个用,分割
}

func (s *service) CreateMenu(ctx context.Context, menuData *CreateMenuData) (err error) {
	qb := admin_menu.NewQueryBuilder()
	qb.WhereAdminId(mysql.EqualPredicate, menuData.AdminId)
	if _, err = qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return
	}

	ActionArr := strings.Split(menuData.Actions, ",")
	for _, v := range ActionArr {
		createModel := admin_menu.NewModel()
		createModel.AdminId = menuData.AdminId
		createModel.MenuId = cast.ToInt64(v)
		createModel.CreatedUser = core.SessionUserInfo(ctx).UserName

		_, err = createModel.Create(mysql.DB().GetDbW().WithContext(ctx))
		if err != nil {
			return
		}
	}

	return
}
