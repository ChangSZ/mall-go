package ums_menu

import (
	"context"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
	"github.com/ChangSZ/mall-go/pkg/log"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Create(ctx context.Context, umsMenu *ums_menu.UmsMenu) (int64, error) {
	data := ums_menu.NewModel()
	data.ParentId = umsMenu.ParentId
	data.Title = umsMenu.Title
	data.Level = umsMenu.Level
	data.Sort = umsMenu.Sort
	data.Name = umsMenu.Name
	data.Icon = umsMenu.Icon
	data.Hidden = umsMenu.Hidden
	s.updateLevel(ctx, data)
	return data.Create(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) updateLevel(ctx context.Context, umsMenu *ums_menu.UmsMenu) {
	if umsMenu.ParentId == 0 {
		umsMenu.Level = 0
	} else {
		qb := ums_menu.NewQueryBuilder()
		qb = qb.WhereId(mysql.EqualPredicate, umsMenu.ParentId)
		parentMenu, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			log.WithTrace(ctx).Errorf("查询父级菜单信息失败, parentId: %v, err: %v", umsMenu.ParentId, err)
		}
		if parentMenu != nil {
			umsMenu.Level = parentMenu.Level + 1
		} else {
			umsMenu.Level = 0
		}
	}
}

func (s *service) Update(ctx context.Context, id int64, umsMenu *ums_menu.UmsMenu) (int64, error) {
	data := ums_menu.NewModel()
	data.Id = id
	data.ParentId = umsMenu.ParentId
	data.Title = umsMenu.Title
	data.Level = umsMenu.Level
	data.Sort = umsMenu.Sort
	data.Name = umsMenu.Name
	data.Icon = umsMenu.Icon
	data.Hidden = umsMenu.Hidden

	s.updateLevel(ctx, data)

	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Update(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) GetItem(ctx context.Context, id int64) (*ums_menu.UmsMenu, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.First(mysql.DB().GetDbR().WithContext(ctx))
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	cnt, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	if cnt == 0 {
		return 0, nil
	}
	return cnt, qb.Delete(mysql.DB().GetDbW().WithContext(ctx))
}

func (s *service) List(ctx context.Context,
	parentId int64, pageSize, pageNum int) ([]*ums_menu.UmsMenu, int64, error) {
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
	return list, count, err
}

type UmsMenuNode struct {
	Id         int64         `json:"id"`
	ParentId   int64         `json:"parentId"`
	CreateTime time.Time     `json:"createTime"`
	Title      string        `json:"title"`
	Level      int32         `json:"level"`
	Sort       int32         `json:"sort"`
	Name       string        `json:"name"`
	Icon       string        `json:"icon"`
	Hidden     int32         `json:"hidden"`
	Children   []UmsMenuNode `json:"children"`
}

func (s *service) TreeList(ctx context.Context) ([]UmsMenuNode, error) {
	qb := ums_menu.NewQueryBuilder()
	menuList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	var result []UmsMenuNode
	for _, menu := range menuList {
		if menu.ParentId == 0 {
			result = append(result, s.covertMenuNode(ctx, menu, menuList))
		}
	}
	return result, nil
}

func (s *service) UpdateHidden(ctx context.Context, id int64, hidden int32) (int64, error) {
	data := ums_menu.NewModel()
	data.Id = id
	data.Hidden = hidden

	qb := ums_menu.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Update(mysql.DB().GetDbW().WithContext(ctx), data)
}

func (s *service) covertMenuNode(ctx context.Context,
	menu *ums_menu.UmsMenu, menuList []*ums_menu.UmsMenu) UmsMenuNode {
	node := UmsMenuNode{}
	node.Id = menu.Id
	node.ParentId = menu.ParentId
	node.CreateTime = menu.CreateTime
	node.Title = menu.Title
	node.Level = menu.Level
	node.Sort = menu.Sort
	node.Name = menu.Name
	node.Icon = menu.Icon
	node.Hidden = menu.Hidden
	node.Children = make([]UmsMenuNode, 0)
	for _, subMenu := range menuList {
		if subMenu.ParentId == menu.Id {
			node.Children = append(node.Children, s.covertMenuNode(ctx, subMenu, menuList))
		}
	}
	return node
}
