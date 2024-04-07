package ums_user

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
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

type umsUserCacheService struct {
	cache   redis.Repo
	service *service
}

func (s *umsUserCacheService) DelAdmin(ctx core.Context, adminId int64) {
	admin, _ := s.service.GetItem(ctx, adminId)
	if admin != nil {
		key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + admin.Username
		s.cache.Del(key)
	}
}

func (s *umsUserCacheService) DelResourceList(ctx core.Context, adminId int64) {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	s.cache.Del(key)
}

func (s *umsUserCacheService) DelResourceListByRole(ctx core.Context, roleId int64) {

}

func (s *umsUserCacheService) DelResourceListByRoleIds(ctx core.Context, roleIds []int64) {

}

func (s *umsUserCacheService) DelResourceListByResource(ctx core.Context, resourceId int64) {

}

func (s *umsUserCacheService) GetAdmin(ctx core.Context, username string) *ums_admin.UmsAdmin {
	key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + username
	ret, err := s.cache.Get(key)
	if err != nil {
		return nil
	}
	var data = &ums_admin.UmsAdmin{}
	if err := json.Unmarshal([]byte(ret), data); err != nil {
		return nil
	}
	return data
}

func (s *umsUserCacheService) SetAdmin(ctx core.Context, admin *ums_admin.UmsAdmin) {
	key := REDIS_DATABASE + ":" + REDIS_KEY_ADMIN + ":" + admin.Username
	adminBytes, _ := json.Marshal(admin)
	s.cache.Set(key, string(adminBytes), REDIS_EXPIRE)
}

func (s *umsUserCacheService) GetResourceList(ctx core.Context, adminId int64) []*ums_resource.UmsResource {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	ret, err := s.cache.Get(key)
	if err != nil {
		return nil
	}
	var data = make([]*ums_resource.UmsResource, 0)
	if err := json.Unmarshal([]byte(ret), &data); err != nil {
		return nil
	}
	return data
}

func (s *umsUserCacheService) SetResourceList(ctx core.Context,
	adminId int64, resourceList []*ums_resource.UmsResource) {
	key := fmt.Sprintf("%s:%s:%d", REDIS_DATABASE, REDIS_KEY_RESOURCE_LIST, adminId)
	resourceListBytes, _ := json.Marshal(resourceList)
	s.cache.Set(key, string(resourceListBytes), REDIS_EXPIRE)
}
