package ums_role

import (
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_menu"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) GetMenuList(ctx core.Context, adminId int64) ([]ums_menu.UmsMenu, error) {
	return dao.GetMenuList(mysql.DB().GetDbR(), adminId)
}
