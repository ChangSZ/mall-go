package ums_user

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role_resource_relation"
)

type service struct {
	cacheService *umsUserCacheService
}

func New() *service {
	return &service{&umsUserCacheService{}}
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
	roleResourceRelations, err := queryBuilder.QueryAll(mysql.DB().GetDbR())
	if err != nil {
		return nil, err
	}
	resourceIds := make([]int64, 0, len(resourceList))
	for _, relation := range roleResourceRelations {
		resourceIds = append(resourceIds, relation.ResourceId)
	}

	resourceQueryBuilder := ums_resource.NewQueryBuilder()
	resourceQueryBuilder = resourceQueryBuilder.WhereIdIn(resourceIds)
	ret, err := resourceQueryBuilder.QueryAll(mysql.DB().GetDbR())
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
	admin, err := queryBuilder.First(mysql.DB().GetDbR())
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
	return queryBuilder.First(mysql.DB().GetDbR())
}
