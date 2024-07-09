package menu

import (
	"context"

	"gorm.io/gorm"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/menu_action"
)

func (s *service) DeleteAction(ctx context.Context, id int64) (err error) {
	// 先查询 id 是否存在
	_, err = menu_action.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereId(mysql.EqualPredicate, id).
		First(mysql.DB().GetDbR().WithContext(ctx))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := menu_action.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	_, err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	return
}
