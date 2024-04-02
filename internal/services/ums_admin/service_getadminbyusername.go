package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
)

// redis:
//
//	database: mall
//	key:
//	  admin: 'ums:admin'
//	  resourceList: 'ums:resourceList'
//	expire:
//	  common: 86400 # 24小时

var (
	REDIS_DATABASE          = "mall"
	REDIS_KEY_ADMIN         = "ums:admin"
	REDIS_KEY_RESOURCE_LIST = "ums:resourceList"
	REDIS_EXPIRE            = 86400
)

func (s *service) GetAdminByUsername(ctx core.Context, username string) (*ums_admin.UmsAdmin, error) {
	return nil, nil
}
