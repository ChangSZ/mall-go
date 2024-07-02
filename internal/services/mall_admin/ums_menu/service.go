package ums_menu

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/pkg/copy"
	"github.com/ChangSZ/mall-go/pkg/log"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, param dto.UmsMenuParam) (int64, error) {
	data := &ums_menu.UmsMenu{}
	copy.AssignStruct(&param, data)
	data.Level = s.getNewLevel(ctx, data.ParentId)
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) getNewLevel(ctx context.Context, parentId int64) int32 {
	if parentId == 0 {
		return 0
	}

	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, parentId)
	parentMenu, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		log.WithTrace(ctx).Errorf("查询父级菜单信息失败, parentId: %v, err: %v", parentId, err)
	}
	if parentMenu != nil {
		return parentMenu.Level + 1
	} else {
		return 0
	}
}

func (s *service) Update(ctx context.Context, id int64, param dto.UmsMenuParam) (int64, error) {
	data := map[string]interface{}{
		"parent_id": param.ParentId,
		"title":     param.Title,
		"level":     s.getNewLevel(ctx, param.ParentId),
		"sort":      param.Sort,
		"name":      param.Name,
		"icon":      param.Icon,
		"hidden":    param.Hidden,
	}
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*dto.UmsMenu, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	data, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}
	res := &dto.UmsMenu{}
	copy.AssignStruct(data, res)
	return res, nil
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) List(ctx context.Context,
	parentId int64, pageSize, pageNum int) ([]dto.UmsMenu, int64, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereParentId(mysql.EqualPredicate, parentId)

	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		OrderBySort(false).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.UmsMenu, 0, len(list))
	for _, v := range list {
		tmp := dto.UmsMenu{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, err
}

func (s *service) TreeList(ctx context.Context) ([]dto.UmsMenuNode, error) {
	qb := ums_menu.NewQueryBuilder()
	menuList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	var result []dto.UmsMenuNode
	for _, menu := range menuList {
		if menu.ParentId == 0 {
			result = append(result, s.covertMenuNode(ctx, menu, menuList))
		}
	}
	return result, nil
}

func (s *service) UpdateHidden(ctx context.Context, id int64, hidden int32) (int64, error) {
	data := map[string]interface{}{
		"hidden": hidden,
	}
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) covertMenuNode(ctx context.Context,
	menu *ums_menu.UmsMenu, menuList []*ums_menu.UmsMenu) dto.UmsMenuNode {
	node := dto.UmsMenuNode{}
	copy.AssignStruct(menu, &node)
	node.Children = make([]dto.UmsMenuNode, 0)
	for _, subMenu := range menuList {
		if subMenu.ParentId == menu.Id {
			node.Children = append(node.Children, s.covertMenuNode(ctx, subMenu, menuList))
		}
	}
	return node
}
