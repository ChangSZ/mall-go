package admin

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/pkg/password"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/admin"
)

type CreateAdminData struct {
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
}

func (s *service) Create(ctx context.Context, adminData *CreateAdminData) (id int32, err error) {
	model := admin.NewModel()
	model.Username = adminData.Username
	model.Password = password.GeneratePassword(adminData.Password)
	model.Nickname = adminData.Nickname
	model.Mobile = adminData.Mobile
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(mysql.DB().GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
