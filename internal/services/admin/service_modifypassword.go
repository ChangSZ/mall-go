package admin

import (
	"context"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/admin"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
)

func (s *service) ModifyPassword(ctx context.Context, id int64, newPassword string) (err error) {
	data := map[string]interface{}{
		"password":     password.GeneratePassword(newPassword),
		"updated_user": core.SessionUserInfo(ctx).UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	_, err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}

	redis.Cache().Del(ctx, configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id))
	return
}
