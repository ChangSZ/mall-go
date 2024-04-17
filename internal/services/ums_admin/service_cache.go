package ums_admin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
)

var (
	REDIS_DATABASE          = configs.Get().Redis.Database
	REDIS_KEY_ADMIN         = configs.Get().Redis.Key.Admin
	REDIS_KEY_RESOURCE_LIST = configs.Get().Redis.Key.ResourceList
	REDIS_EXPIRE            = time.Duration(configs.Get().Redis.Expire.Common) * time.Second
)

type umsAdminCacheService struct {
	service *service
}

func (s *umsAdminCacheService) DelAdmin(ctx context.Context, adminId int64) {
	admin, _ := s.service.GetItem(ctx, adminId)
	if admin != nil {
		key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + admin.Username
		redis.Cache().Del(ctx, key)
	}
}

func (s *umsAdminCacheService) DelResourceList(ctx context.Context, adminId int64) {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	redis.Cache().Del(ctx, key)
}

func (s *umsAdminCacheService) DelResourceListByRole(ctx context.Context, roleId int64) {

}

func (s *umsAdminCacheService) DelResourceListByRoleIds(ctx context.Context, roleIds []int64) {

}

func (s *umsAdminCacheService) DelResourceListByResource(ctx context.Context, resourceId int64) {

}

func (s *umsAdminCacheService) GetAdmin(ctx context.Context, username string) *ums_admin.UmsAdmin {
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

func (s *umsAdminCacheService) SetAdmin(ctx context.Context, admin *ums_admin.UmsAdmin) {
	key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + admin.Username
	adminBytes, _ := json.Marshal(admin)
	redis.Cache().Set(ctx, key, string(adminBytes), REDIS_EXPIRE)
}

func (s *umsAdminCacheService) GetResourceList(ctx context.Context, adminId int64) []*ums_resource.UmsResource {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	ret, err := redis.Cache().Get(ctx, key)
	if err != nil {
		return nil
	}
	var data = make([]*ums_resource.UmsResource, 0)
	if err := json.Unmarshal([]byte(ret), &data); err != nil {
		return nil
	}
	return data
}

func (s *umsAdminCacheService) SetResourceList(ctx context.Context,
	adminId int64, resourceList []*ums_resource.UmsResource) {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	resourceListBytes, _ := json.Marshal(resourceList)
	redis.Cache().Set(ctx, key, string(resourceListBytes), REDIS_EXPIRE)
}
