package ums_role

import (
	"context"
	"strings"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role_menu_relation"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role_resource_relation"
	"github.com/ChangSZ/mall-go/internal/services/ums_admin"
)

type service struct {
	adminCacheService *ums_admin.UmsAdminCacheService
}

func New() Service {
	return &service{ums_admin.NewCacheService()}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, umsRole *ums_role.UmsRole) (int64, error) {
	return umsRole.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, umsRole *ums_role.UmsRole) (int64, error) {
	umsRole.Id = id
	qb := ums_role.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Update(mysql.DB().GetDbW().WithContext(ctx), umsRole)
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil || cnt == 0 {
		return 0, err
	}
	if err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByRoleIds(ctx, ids)
	return cnt, nil
}

func (s *service) ListAll(ctx context.Context) ([]*ums_role.UmsRole, error) {
	qb := ums_role.NewQueryBuilder()
	return qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
}

func (s *service) List(ctx context.Context, keyword string, pageSize, pageNum int) ([]*ums_role.UmsRole, int64, error) {
	qb := ums_role.NewQueryBuilder()
	if strings.TrimSpace(keyword) != "" {
		qb = qb.WhereName(mysql.LikePredicate, "%"+keyword+"%")
	}
	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	return list, count, err
}

func (s *service) GetMenuList(ctx context.Context, adminId int64) ([]ums_menu.UmsMenu, error) {
	return new(dao.UmsRoleDao).GetMenuList(mysql.DB().GetDbR(), adminId)
}

func (s *service) ListMenu(ctx context.Context, roleId int64) ([]ums_menu.UmsMenu, error) {
	return new(dao.UmsRoleDao).GetMenuListByRoleId(mysql.DB().GetDbR().WithContext(ctx), roleId)
}

func (s *service) ListResource(ctx context.Context, roleId int64) ([]ums_resource.UmsResource, error) {
	return new(dao.UmsRoleDao).GetResourceListByRoleId(mysql.DB().GetDbR().WithContext(ctx), roleId)
}

func (s *service) AllocMenu(ctx context.Context, roleId int64, menuIds []int64) (int64, error) {
	// 先删除原有关系
	qb := ums_role_menu_relation.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, roleId)
	if err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, err
	}

	// 批量插入新关系
	relationList := make([]*ums_role_menu_relation.UmsRoleMenuRelation, 0, len(menuIds))
	for _, menuId := range menuIds {
		relation := ums_role_menu_relation.NewModel()
		relation.RoleId = roleId
		relation.MenuId = menuId
		relationList = append(relationList, relation)
	}
	if err := mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(relationList, len(menuIds)).Error; err != nil {
		return 0, err
	}
	return int64(len(menuIds)), nil
}

func (s *service) AllocResource(ctx context.Context, roleId int64, resourceIds []int64) (int64, error) {
	// 先删除原有关系
	qb := ums_role_resource_relation.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, roleId)
	if err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, err
	}

	// 批量插入新关系
	relationList := make([]*ums_role_resource_relation.UmsRoleResourceRelation, 0, len(resourceIds))
	for _, resourceId := range resourceIds {
		relation := ums_role_resource_relation.NewModel()
		relation.RoleId = roleId
		relation.ResourceId = resourceId
		relationList = append(relationList, relation)
	}
	if err := mysql.DB().GetDbW().WithContext(ctx).CreateInBatches(relationList, len(resourceIds)).Error; err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByRole(ctx, roleId)
	return int64(len(resourceIds)), nil
}
