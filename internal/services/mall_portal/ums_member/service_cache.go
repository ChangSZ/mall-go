package ums_member

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
)

var (
	REDIS_DATABASE         = configs.Get().Redis.Database
	REDIS_EXPIRE           = time.Duration(configs.Get().Redis.Expire.Common) * time.Second
	REDIS_EXPIRE_AUTH_CODE = time.Duration(configs.Get().Redis.Expire.AuthCode) * time.Second
	REDIS_KEY_MEMBER       = configs.Get().Redis.Key.Member
	REDIS_KEY_AUTH_CODE    = configs.Get().Redis.Key.AuthCode
)

type UmsMemberCacheService struct{}

func NewCacheService() *UmsMemberCacheService {
	return &UmsMemberCacheService{}
}

func (s *UmsMemberCacheService) DelMember(ctx context.Context, memberId int64) {
	member, _ := New().GetById(ctx, memberId)
	if member != nil {
		key := REDIS_DATABASE + ":" + REDIS_KEY_MEMBER + ":" + member.Username
		redis.Cache().Del(ctx, key)
	}
}

func (s *UmsMemberCacheService) GetMember(ctx context.Context, username string) *ums_member.UmsMember {
	key := REDIS_DATABASE + ":" + REDIS_KEY_MEMBER + ":" + username
	ret, err := redis.Cache().Get(ctx, key)
	if err != nil {
		return nil
	}
	var data = &ums_member.UmsMember{}
	if err := json.Unmarshal([]byte(ret), data); err != nil {
		return nil
	}
	return data
}

func (s *UmsMemberCacheService) SetMember(ctx context.Context, member *ums_member.UmsMember) {
	key := REDIS_DATABASE + ":" + REDIS_KEY_MEMBER + ":" + member.Username
	memberBytes, _ := json.Marshal(member)
	redis.Cache().Set(ctx, key, string(memberBytes), REDIS_EXPIRE)
}

func (s *UmsMemberCacheService) SetAuthCode(ctx context.Context, telephone, authCode string) {
	key := REDIS_DATABASE + ":" + REDIS_KEY_AUTH_CODE + ":" + telephone
	redis.Cache().Set(ctx, key, authCode, REDIS_EXPIRE_AUTH_CODE)
}

func (s *UmsMemberCacheService) GetAuthCode(ctx context.Context, telephone string) string {
	key := REDIS_DATABASE + ":" + REDIS_KEY_AUTH_CODE + ":" + telephone
	ret, _ := redis.Cache().Get(ctx, key)
	return ret
}
