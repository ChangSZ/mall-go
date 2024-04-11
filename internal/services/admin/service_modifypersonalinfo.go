package admin

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/admin"
)

type ModifyData struct {
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) ModifyPersonalInfo(ctx context.Context, id int32, modifyData *ModifyData) (err error) {
	data := map[string]interface{}{
		"nickname":     modifyData.Nickname,
		"mobile":       modifyData.Mobile,
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	return
}
