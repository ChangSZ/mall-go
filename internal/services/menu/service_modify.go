package menu

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/menu"
)

type UpdateMenuData struct {
	Name string // 菜单名称
	Link string // 链接地址
	Icon string // 图标
}

func (s *service) Modify(ctx context.Context, id int64, menuData *UpdateMenuData) (err error) {
	data := map[string]interface{}{
		"name":         menuData.Name,
		"link":         menuData.Link,
		"icon":         menuData.Icon,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := menu.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	_, err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	return
}
