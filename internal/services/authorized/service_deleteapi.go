package authorized

import (
	"context"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/authorized_api"
	"github.com/ChangSZ/mall-go/internal/repository/redis"

	"gorm.io/gorm"
)

func (s *service) DeleteAPI(ctx context.Context, id int32) (err error) {
	// 先查询 id 是否存在
	authorizedApiInfo, err := authorized_api.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereId(mysql.EqualPredicate, id).
		First(mysql.DB().GetDbR().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := authorized_api.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(mysql.DB().GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	redis.Cache().Del(configs.RedisKeyPrefixSignature+authorizedApiInfo.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}
