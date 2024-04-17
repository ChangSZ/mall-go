package ums_admin

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role_resource_relation"
	"github.com/ChangSZ/mall-go/pkg/jwt"
	"github.com/ChangSZ/mall-go/pkg/password"
)

var (
	jwtConfig    = configs.Get().Jwt
	jwtTokenUtil = jwt.NewJwtTokenUtil(jwtConfig.Secret, jwtConfig.Expiration, jwtConfig.TokenHead)
)

type service struct{ cacheService *umsAdminCacheService }

func New() Service {
	return &service{&umsAdminCacheService{}}
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

func (s *service) Register(ctx context.Context, umsAdminParam *UmsAdminParam) (*ums_admin.UmsAdmin, error) {
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
	umsAdminList, err := queryBuilder.QueryAll(mysql.DB().GetDbR())
	if err != nil {
		return nil, err
	}
	if len(umsAdminList) > 0 {
		return nil, fmt.Errorf("用户名已存在")
	}

	_, err = umsAdmin.Create(mysql.DB().GetDbW().WithContext(ctx))
	return umsAdmin, err
}

func (s *service) Login(ctx context.Context, username, passwd string) (string, error) {
	var token string
	userDetails, err := s.LoadUserByUsername(ctx, username)
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

func (s *service) RefreshToken(ctx context.Context, oldToken string) (string, error) {
	return jwtTokenUtil.RefreshHeadToken(oldToken, 1800) // 30 minutes
}

type UpdateAdminPasswordParam struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func (s *service) UpdatePassword(ctx context.Context, updatePasswordParam *UpdateAdminPasswordParam) (int64, error) {
	return 0, nil
}

func (s *service) GetRoleList(ctx context.Context, adminId int64) ([]ums_role.UmsRole, error) {
	return dao.GetRoleList(mysql.DB().GetDbR(), adminId)
}

func (s *service) List(ctx context.Context, keyword string, pageSize, pageNum int) ([]ums_admin.UmsAdmin, int64, error) {
	return dao.AdminPageList(ctx, mysql.DB().GetDbR(), keyword, pageSize, pageNum)
}

func (s *service) GetResourceList(ctx context.Context, adminId int64) ([]*ums_resource.UmsResource, error) {
	// 先从缓存中获取数据
	resourceList := s.cacheService.GetResourceList(ctx, adminId)
	if len(resourceList) != 0 {
		return resourceList, nil
	}

	// 缓存中没有从数据库中获取
	queryBuilder := ums_role_resource_relation.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereRoleId(mysql.EqualPredicate, adminId)
	roleResourceRelations, err := queryBuilder.QueryAll(mysql.DB().GetDbR())
	if err != nil {
		return nil, err
	}
	resourceIds := make([]int64, 0, len(resourceList))
	for _, relation := range roleResourceRelations {
		resourceIds = append(resourceIds, relation.ResourceId)
	}

	resourceQueryBuilder := ums_resource.NewQueryBuilder()
	resourceQueryBuilder = resourceQueryBuilder.WhereIdIn(resourceIds)
	ret, err := resourceQueryBuilder.QueryAll(mysql.DB().GetDbR())
	if len(ret) != 0 {
		// 将数据库中的数据存入缓存中
		s.cacheService.SetResourceList(ctx, adminId, ret)
	}
	return ret, err
}

func (s *service) GetAdminByUsername(ctx context.Context, username string) (*ums_admin.UmsAdmin, error) {
	// 先从缓存中获取数据
	admin := s.cacheService.GetAdmin(ctx, username)
	if admin != nil {
		return admin, nil
	}

	// 缓存中没有从数据库中获取
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder.WhereUsername(mysql.EqualPredicate, username)
	admin, err := queryBuilder.First(mysql.DB().GetDbR())
	if err != nil {
		return nil, err
	}

	// 将数据库中的数据存入缓存中
	s.cacheService.SetAdmin(ctx, admin)
	return admin, nil
}

func (s *service) LoadUserByUsername(ctx context.Context, username string) (*AdminUserDetails, error) {
	// 获取用户信息
	admin, err := s.GetAdminByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if admin != nil {
		resourceList, err := s.GetResourceList(ctx, admin.Id)
		if err != nil {
			return nil, err
		}
		return &AdminUserDetails{admin, resourceList}, nil
	}
	return nil, fmt.Errorf("用户名或密码错误")
}

func (s *service) GetItem(ctx context.Context, id int64) (*ums_admin.UmsAdmin, error) {
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereId(mysql.EqualPredicate, id)
	return queryBuilder.First(mysql.DB().GetDbR())
}

func (s *service) Update(ctx context.Context, id int64, admin *ums_admin.UmsAdmin) (int64, error) {
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereId(mysql.EqualPredicate, id)
	return queryBuilder.Update(mysql.DB().GetDbW().WithContext(ctx), admin)
}
