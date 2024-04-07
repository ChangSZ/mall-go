package ums_admin

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
	"github.com/ChangSZ/mall-go/internal/services/ums_user"
	"github.com/ChangSZ/mall-go/pkg/jwt"
	"github.com/ChangSZ/mall-go/pkg/password"
)

var (
	jwtConfig    = configs.Get().Jwt
	jwtTokenUtil = jwt.NewJwtTokenUtil(jwtConfig.Secret, jwtConfig.Expiration, jwtConfig.TokenHead)
)

type service struct {
	db mysql.Repo
}

func New(db mysql.Repo) Service {
	s := &service{db: db}
	return s
}

func (s *service) i() {}

type UmsAdminParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Icon     string `json:"icon"`
	Email    string `json:"email" binding:"email"`
	NickName string `json:"nickName"`
	Note     string `json:"note"`
}

func (s *service) Register(ctx core.Context, umsAdminParam *UmsAdminParam) (*ums_admin.UmsAdmin, error) {
	umsAdmin := ums_admin.NewModel()
	umsAdmin.Username = umsAdminParam.Username
	encodePassword, err := password.Encoder.Encode(umsAdminParam.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to encode password: %w", err)
	}
	umsAdmin.Password = encodePassword
	umsAdmin.Icon = umsAdminParam.Icon
	umsAdmin.Email = umsAdminParam.Email
	umsAdmin.NickName = umsAdminParam.NickName
	umsAdmin.Note = umsAdminParam.Note
	umsAdmin.LoginTime = time.Now()
	umsAdmin.Status = 1

	// 查询是否有相同用户名的用户
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder.WhereUsername(mysql.EqualPredicate, umsAdmin.Username)
	umsAdminList, err := queryBuilder.QueryAll(s.db.GetDbR())
	if err != nil {
		return nil, err
	}
	if len(umsAdminList) > 0 {
		return nil, fmt.Errorf("用户名已存在")
	}

	_, err = umsAdmin.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	return umsAdmin, err
}

func (s *service) Login(ctx core.Context, username, passwd string) (string, error) {
	var token string
	userDetails, err := ums_user.DefalutService.LoadUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if !password.Encoder.Matches(passwd, userDetails.GetPassword()) {
		return "", fmt.Errorf("密码不正确")
	}
	if !userDetails.IsEnabled() {
		return "", fmt.Errorf("账号已被禁用")
	}

	token, err = jwtTokenUtil.GenerateToken(username)
	if err != nil {
		return "", fmt.Errorf("生成token失败: %w", err)
	}
	return token, nil
}

func (s *service) RefreshToken(ctx core.Context, oldToken string) (string, error) {
	return jwtTokenUtil.RefreshHeadToken(oldToken, 1800) // 30 minutes
}

type UpdateAdminPasswordParam struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func (s *service) UpdatePassword(ctx core.Context, updatePasswordParam *UpdateAdminPasswordParam) (int64, error) {
	return 0, nil
}

func (s *service) GetRoleList(ctx core.Context, adminId int64) ([]ums_role.UmsRole, error) {
	return dao.GetRoleList(s.db.GetDbR(), adminId)
}

// func (s *service) GetCacheService() *umsAdminCacheService {
// 	return s.cache
// }
