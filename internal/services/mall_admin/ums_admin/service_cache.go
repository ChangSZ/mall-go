package ums_admin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ChangSZ/golib/log"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin_role_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
)

var (
	REDIS_DATABASE          = configs.Get().Redis.Database
	REDIS_KEY_ADMIN         = configs.Get().Redis.Key.Admin
	REDIS_KEY_RESOURCE_LIST = configs.Get().Redis.Key.ResourceList
	REDIS_EXPIRE            = time.Duration(configs.Get().Redis.Expire.Common) * time.Second
)

type UmsAdminCacheService struct{}

func NewCacheService() *UmsAdminCacheService {
	return &UmsAdminCacheService{}
}

func (s *UmsAdminCacheService) DelAdmin(ctx context.Context, adminId int64) {
	admin, _ := New().GetItem(ctx, adminId)
	if admin != nil {
		key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + admin.Username
		redis.Cache().Del(ctx, key)
	}
}

func (s *UmsAdminCacheService) DelResourceList(ctx context.Context, adminId int64) {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	redis.Cache().Del(ctx, key)
}

func (s *UmsAdminCacheService) DelResourceListByRole(ctx context.Context, roleId int64) {
	s.DelResourceListByRoleIds(ctx, []int64{roleId})
}

func (s *UmsAdminCacheService) DelResourceListByRoleIds(ctx context.Context, roleIds []int64) {
	qb := ums_admin_role_relation.NewQueryBuilder()
	qb = qb.WhereRoleIdIn(roleIds)
	relationList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		log.WithTrace(ctx).Errorf("获取UmsAdminRoleRelationList失败, roleIds: %v, err: %v", roleIds, err)
		return
	}
	if len(relationList) != 0 {
		keyPrefix := REDIS_DATABASE + ":" + REDIS_KEY_RESOURCE_LIST + ":"
		for _, v := range relationList {
			key := fmt.Sprintf("%s%d", keyPrefix, v.AdminId)
			redis.Cache().Del(ctx, key)
		}
	}
}

func (s *UmsAdminCacheService) DelResourceListByResource(ctx context.Context, resourceId int64) {
	adminIdList, err := new(dao.UmsAdminRoleRelationDao).GetAdminIdList(mysql.DB().GetDbW().WithContext(ctx), resourceId)
	if err != nil {
		log.WithTrace(ctx).Errorf("获取AdminIdList失败, resourceId: %v, err: %v", resourceId, err)
		return
	}
	if len(adminIdList) != 0 {
		keyPrefix := REDIS_DATABASE + ":" + REDIS_KEY_RESOURCE_LIST + ":"
		for _, v := range adminIdList {
			key := fmt.Sprintf("%s%d", keyPrefix, v)
			redis.Cache().Del(ctx, key)
		}
	}
}

func (s *UmsAdminCacheService) GetAdmin(ctx context.Context, username string) *ums_admin.UmsAdmin {
	key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + username
	ret, err := redis.Cache().Get(ctx, key)
	if err != nil {
		return nil
	}
	var data = &ums_admin.UmsAdmin{}
	if err := json.Unmarshal([]byte(ret), data); err != nil {
		return nil
	}
	return data
}

func (s *UmsAdminCacheService) SetAdmin(ctx context.Context, admin *ums_admin.UmsAdmin) {
	key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + admin.Username
	adminBytes, _ := json.Marshal(admin)
	redis.Cache().Set(ctx, key, string(adminBytes), REDIS_EXPIRE)
}

func (s *UmsAdminCacheService) GetResourceList(ctx context.Context, adminId int64) []ums_resource.UmsResource {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	ret, err := redis.Cache().Get(ctx, key)
	if err != nil {
		return nil
	}
	var data = make([]ums_resource.UmsResource, 0)
	if err := json.Unmarshal([]byte(ret), &data); err != nil {
		return nil
	}
	return data
}

func (s *UmsAdminCacheService) SetResourceList(ctx context.Context,
	adminId int64, resourceList []ums_resource.UmsResource) {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	resourceListBytes, _ := json.Marshal(resourceList)
	redis.Cache().Set(ctx, key, string(resourceListBytes), REDIS_EXPIRE)
}
