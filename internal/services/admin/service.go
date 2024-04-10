package admin

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/admin"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx context.Context, adminData *CreateAdminData) (id int32, err error)
	PageList(ctx context.Context, searchData *SearchData) (listData []*admin.Admin, err error)
	PageListCount(ctx context.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx context.Context, id int32, used int32) (err error)
	Delete(ctx context.Context, id int32) (err error)
	Detail(ctx context.Context, searchOneData *SearchOneData) (info *admin.Admin, err error)
	ResetPassword(ctx context.Context, id int32) (err error)
	ModifyPassword(ctx context.Context, id int32, newPassword string) (err error)
	ModifyPersonalInfo(ctx context.Context, id int32, modifyData *ModifyData) (err error)

	CreateMenu(ctx context.Context, menuData *CreateMenuData) (err error)
	ListMenu(ctx context.Context, searchData *SearchListMenuData) (menuData []ListMenuData, err error)
	MyMenu(ctx context.Context, searchData *SearchMyMenuData) (menuData []ListMyMenuData, err error)
	MyAction(ctx context.Context, searchData *SearchMyActionData) (actionData []MyActionData, err error)
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
