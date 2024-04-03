package ums_admin

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_role_resource_relation"
	"github.com/ChangSZ/mall-go/internal/repository/redis"
	"github.com/ChangSZ/mall-go/pkg/jwt"
	"github.com/ChangSZ/mall-go/pkg/password"
)

type service struct {
	db    mysql.Repo
	cache *umsAdminCacheService
}

func New(db mysql.Repo, cache redis.Repo) Service {
	s := &service{db: db}
	s.cache = &umsAdminCacheService{cache: cache, service: s} // 为了在cache中使用service的一些函数
	return s
}

func (s *service) i() {}

func (s *service) GetAdminByUsername(ctx core.Context, username string) (*ums_admin.UmsAdmin, error) {
	// 先从缓存中获取数据
	admin := s.cache.GetAdmin(ctx, username)
	if admin != nil {
		return admin, nil
	}

	// 缓存中没有从数据库中获取
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder.WhereUsername(mysql.EqualPredicate, username)
	admin, err := queryBuilder.First(s.db.GetDbR())
	if err != nil {
		return nil, err
	}

	// 将数据库中的数据存入缓存中
	s.cache.SetAdmin(ctx, admin)
	return admin, nil
}

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
	config := configs.Get().Jwt
	jwtTokenUtil := jwt.NewJwtTokenUtil(config.Secret, config.Expiration, config.TokenHead)
	token, err = jwtTokenUtil.GenerateToken(username)
	if err != nil {
		return "", fmt.Errorf("生成token失败: %w", err)
	}
	return token, nil
}

type UpdateAdminPasswordParam struct {
	Username    string `json:"username" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func (s *service) UpdatePassword(ctx core.Context, updatePasswordParam *UpdateAdminPasswordParam) (int64, error) {
	return 0, nil
}

func (s *service) GetItem(ctx core.Context, id int64) (*ums_admin.UmsAdmin, error) {
	queryBuilder := ums_admin.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereId(mysql.EqualPredicate, id)
	return queryBuilder.First(s.db.GetDbR())
}

func (s *service) GetResourceList(ctx core.Context, adminId int64) ([]*ums_resource.UmsResource, error) {
	// 先从缓存中获取数据
	resourceList := s.cache.GetResourceList(ctx, adminId)
	if len(resourceList) != 0 {
		return resourceList, nil
	}

	// 缓存中没有从数据库中获取
	queryBuilder := ums_role_resource_relation.NewQueryBuilder()
	queryBuilder = queryBuilder.WhereRoleId(mysql.EqualPredicate, adminId)
	roleResourceRelations, err := queryBuilder.QueryAll(s.db.GetDbR())
	if err != nil {
		return nil, err
	}
	resourceIds := make([]int64, 0, len(resourceList))
	for _, relation := range roleResourceRelations {
		resourceIds = append(resourceIds, relation.ResourceId)
	}

	resourceQueryBuilder := ums_resource.NewQueryBuilder()
	resourceQueryBuilder = resourceQueryBuilder.WhereIdIn(resourceIds)
	ret, err := resourceQueryBuilder.QueryAll(s.db.GetDbR())
	if len(ret) != 0 {
		// 将数据库中的数据存入缓存中
		s.cache.SetResourceList(ctx, adminId, ret)
	}
	return ret, err
}

func (s *service) LoadUserByUsername(ctx core.Context, username string) (*AdminUserDetails, error) {
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

// func (s *service) GetCacheService() *umsAdminCacheService {
// 	return s.cache
// }
