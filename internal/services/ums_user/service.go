package ums_user

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role_resource_relation"
)

/*
由于依赖的问题, 该包弄成一个公共的服务, 可以供其他包（比如中间件）也能够查询用户信息
*/

var DefalutService *service

type service struct {
	db           mysql.Repo
	cacheService *umsUserCacheService
}

func DefaultService(db mysql.Repo) {
	s := &service{db: db}
	s.cacheService = &umsUserCacheService{service: s}
	DefalutService = s
}

func NewService(db mysql.Repo, cache *umsUserCacheService) *service {
	return &service{db, cache}
}

func (s *service) GetResourceList(ctx core.Context, adminId int64) ([]*ums_resource.UmsResource, error) {
	// 先从缓存中获取数据
	resourceList := s.cacheService.GetResourceList(ctx, adminId)
	if len(resourceList) != 0 {
		return resourceList, nil
	}

	// 缓存中没有从数据库中获取
	queryBuilder := ums_role_resource_relation.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereRoleId(mysql.EqualPredicate, adminId)
	roleResourceRelations, err := queryBuilder.QueryAll(s.db.GetDbR())
	if err != nil {
		return nil, err
	}
	resourceIds := make([]int64, 0, len(resourceList))
	for _, relation := range roleResourceRelations {
		resourceIds = append(resourceIds, relation.ResourceId)
	}

	resourceQueryBuilder := ums_resource.NewQueryBuilder()
	resourceQueryBuilder = resourceQueryBuilder.WhereIdIn(resourceIds)
	ret, err := resourceQueryBuilder.QueryAll(s.db.GetDbR())
	if len(ret) != 0 {
		// 将数据库中的数据存入缓存中
		s.cacheService.SetResourceList(ctx, adminId, ret)
	}
	return ret, err
}

func (s *service) GetAdminByUsername(ctx core.Context, username string) (*ums_admin.UmsAdmin, error) {
	// 先从缓存中获取数据
	admin := s.cacheService.GetAdmin(ctx, username)
	if admin != nil {
		return admin, nil
	}

	// 缓存中没有从数据库中获取
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder.WhereUsername(mysql.EqualPredicate, username)
	admin, err := queryBuilder.First(s.db.GetDbR())
	if err != nil {
		return nil, err
	}

	// 将数据库中的数据存入缓存中
	s.cacheService.SetAdmin(ctx, admin)
	return admin, nil
}

func (s *service) LoadUserByUsername(ctx core.Context, username string) (*AdminUserDetails, error) {
	// 获取用户信息
	admin, err := s.GetAdminByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if admin != nil {
		resourceList, err := s.GetResourceList(ctx, admin.Id)
		if err != nil {
			return nil, err
		}
		return &AdminUserDetails{admin, resourceList}, nil
	}
	return nil, fmt.Errorf("用户名或密码错误")
}

func (s *service) GetItem(ctx core.Context, id int64) (*ums_admin.UmsAdmin, error) {
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereId(mysql.EqualPredicate, id)
	return queryBuilder.First(s.db.GetDbR())
}
