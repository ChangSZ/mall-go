package admin

import (
	"context"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/admin"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
)

func (s *service) UpdateUsed(ctx context.Context, id int32, used int32) (err error) {
	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	redis.Cache().Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}
