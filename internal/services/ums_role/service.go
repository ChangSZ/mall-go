package ums_role

import (
	"context"
	"strings"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
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

func (s *service) Create(ctx context.Context, param dto.UmsRoleParam) (int64, error) {
	data := &ums_role.UmsRole{
		Name:        param.Name,
		Description: param.Description,
		Status:      param.Status,
	}
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) Update(ctx context.Context, id int64, param dto.UmsRoleParam) (int64, error) {
	data := map[string]interface{}{
		"name":        param.Name,
		"description": param.Description,
		"admin_count": param.AdminCount,
		"status":      param.Status,
		"sort":        param.Sort,
	}
	qb := ums_role.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) UpdateStatus(ctx context.Context, id int64, status int32) (int64, error) {
	data := map[string]interface{}{
		"status": status,
	}
	qb := ums_role.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) Delete(ctx context.Context, ids []int64) (int64, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereIdIn(ids)
	cnt, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	s.adminCacheService.DelResourceListByRoleIds(ctx, ids)
	return cnt, nil
}

func (s *service) ListAll(ctx context.Context) ([]dto.UmsRole, error) {
	qb := ums_role.NewQueryBuilder()
	list, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	listData := make([]dto.UmsRole, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsRole{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			AdminCount:  v.AdminCount,
			CreateTime:  v.CreateTime,
			Status:      v.Status,
			Sort:        v.Sort,
		})
	}
	return listData, nil
}

func (s *service) List(ctx context.Context, keyword string, pageSize, pageNum int) ([]dto.UmsRole, int64, error) {
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
	if err != nil {
		return nil, 0, err
	}
	listData := make([]dto.UmsRole, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsRole{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			AdminCount:  v.AdminCount,
			CreateTime:  v.CreateTime,
			Status:      v.Status,
			Sort:        v.Sort,
		})
	}
	return listData, count, err
}

func (s *service) GetMenuList(ctx context.Context, adminId int64) ([]dto.UmsMenu, error) {
	list, err := new(dao.UmsRoleDao).GetMenuList(mysql.DB().GetDbR(), adminId)
	if err != nil {
		return nil, err
	}
	listData := make([]dto.UmsMenu, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsMenu{
			Id:         v.Id,
			ParentId:   v.ParentId,
			CreateTime: v.CreateTime,
			Title:      v.Title,
			Level:      v.Level,
			Sort:       v.Sort,
			Name:       v.Name,
			Icon:       v.Icon,
			Hidden:     v.Hidden,
		})
	}
	return listData, nil
}

func (s *service) ListMenu(ctx context.Context, roleId int64) ([]dto.UmsMenu, error) {
	list, err := new(dao.UmsRoleDao).GetMenuListByRoleId(mysql.DB().GetDbR().WithContext(ctx), roleId)
	if err != nil {
		return nil, err
	}
	listData := make([]dto.UmsMenu, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsMenu{
			Id:         v.Id,
			ParentId:   v.ParentId,
			CreateTime: v.CreateTime,
			Title:      v.Title,
			Level:      v.Level,
			Sort:       v.Sort,
			Name:       v.Name,
			Icon:       v.Icon,
			Hidden:     v.Hidden,
		})
	}
	return listData, nil
}

func (s *service) ListResource(ctx context.Context, roleId int64) ([]dto.UmsResource, error) {
	list, err := new(dao.UmsRoleDao).GetResourceListByRoleId(mysql.DB().GetDbR().WithContext(ctx), roleId)
	if err != nil {
		return nil, err
	}
	listData := make([]dto.UmsResource, 0, len(list))
	for _, v := range list {
		listData = append(listData, dto.UmsResource{
			Id:          v.Id,
			CreateTime:  v.CreateTime,
			Name:        v.Name,
			Url:         v.Url,
			Description: v.Description,
			CategoryId:  v.CategoryId,
		})
	}
	return listData, nil
}

func (s *service) AllocMenu(ctx context.Context, roleId int64, menuIds []int64) (int64, error) {
	// 先删除原有关系
	qb := ums_role_menu_relation.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, roleId)
	if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
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
	if _, err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
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
