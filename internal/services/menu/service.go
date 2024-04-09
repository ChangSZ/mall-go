package menu

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/menu"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/menu_action"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, menuData *CreateMenuData) (id int32, err error)
	Modify(ctx core.Context, id int32, menuData *UpdateMenuData) (err error)
	List(ctx core.Context, searchData *SearchData) (listData []*menu.Menu, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	UpdateSort(ctx core.Context, id int32, sort int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *menu.Menu, err error)

	CreateAction(ctx core.Context, menuActionData *CreateMenuActionData) (id int32, err error)
	ListAction(ctx core.Context, searchListActionData *SearchListActionData) (listData []*menu_action.MenuAction, err error)
	DeleteAction(ctx core.Context, id int32) (err error)
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
