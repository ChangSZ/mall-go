package ums_role

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) GetMenuList(ctx context.Context, adminId int64) ([]ums_menu.UmsMenu, error) {
	return new(dao.UmsRoleDao).GetMenuList(mysql.DB().GetDbR(), adminId)
}
