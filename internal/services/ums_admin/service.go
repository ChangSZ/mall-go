package ums_admin

import (
	"context"
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin_role_relation"
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
		return nil, fmt.Errorf("failed to encode password: %v", err)
	}
	umsAdmin.Password = encodePassword
	umsAdmin.Icon = umsAdminParam.Icon
	umsAdmin.Email = umsAdminParam.Email
	umsAdmin.NickName = umsAdminParam.NickName
	umsAdmin.Note = umsAdminParam.Note
	umsAdmin.LoginTime = time.Now()
	umsAdmin.Status = 1

	// 查询是否有相同用户名的用户
	qb := ums_admin.NewQueryBuilder()
	qb.WhereUsername(mysql.EqualPredicate, umsAdmin.Username)
	umsAdminList, err := qb.QueryAll(mysql.DB().GetDbR())
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
		return "", fmt.Errorf("生成token失败: %v", err)
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

func (s *service) UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) (int64, error) {
	qb := ums_admin.NewQueryBuilder()
	qb = qb.WhereUsername(mysql.EqualPredicate, username)
	adminList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return 0, err
	}
	if len(adminList) == 0 {
		return 0, fmt.Errorf("找不到该用户")
	}
	umsAdmin := adminList[0] // 理论上username应该是唯一的

	if !password.Encoder.Matches(oldPassword, umsAdmin.Password) {
		return 0, fmt.Errorf("旧密码错误")
	}

	newPasswd, err := password.Encoder.Encode(newPassword)
	if err != nil {
		return 0, err
	}
	data := ums_admin.NewModel()
	data.Password = newPasswd
	qb = ums_admin.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, umsAdmin.Id)
	cnt, err := qb.Update(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return 0, err
	}
	s.cacheService.DelAdmin(ctx, umsAdmin.Id)
	return cnt, nil
}

func (s *service) GetRoleList(ctx context.Context, adminId int64) ([]ums_role.UmsRole, error) {
	roleRelationDao := new(dao.UmsAdminRoleRelationDao)
	return roleRelationDao.GetRoleList(mysql.DB().GetDbR(), adminId)
}

func (s *service) List(ctx context.Context, keyword string, pageSize, pageNum int) ([]ums_admin.UmsAdmin, int64, error) {
	return new(dao.UmsAdminDao).AdminPageList(ctx, mysql.DB().GetDbR(), keyword, pageSize, pageNum)
}

func (s *service) GetResourceList(ctx context.Context, adminId int64) ([]*ums_resource.UmsResource, error) {
	// 先从缓存中获取数据
	resourceList := s.cacheService.GetResourceList(ctx, adminId)
	if len(resourceList) != 0 {
		return resourceList, nil
	}

	// 缓存中没有从数据库中获取
	qb := ums_role_resource_relation.NewQueryBuilder()
	qb = qb.WhereRoleId(mysql.EqualPredicate, adminId)
	roleResourceRelations, err := qb.QueryAll(mysql.DB().GetDbR())
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
	qb := ums_admin.NewQueryBuilder()
	qb.WhereUsername(mysql.EqualPredicate, username)
	admin, err := qb.First(mysql.DB().GetDbR())
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
	qb := ums_admin.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.First(mysql.DB().GetDbR())
}

func (s *service) Update(ctx context.Context, id int64, admin *ums_admin.UmsAdmin) (int64, error) {
	qb := ums_admin.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.Update(mysql.DB().GetDbW().WithContext(ctx), admin)
}

func (s *service) Delete(ctx context.Context, id int64) (int64, error) {
	qb := ums_admin.NewQueryBuilder()
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

func (s *service) UpdateRole(ctx context.Context, adminId int64, roleIds []int64) (int64, error) {
	if len(roleIds) == 0 {
		return 0, nil
	}
	var count = int64(len(roleIds))

	// 先删除原来的关系
	qb := ums_admin_role_relation.NewQueryBuilder()
	qb = qb.WhereAdminId(mysql.EqualPredicate, adminId)
	if err := qb.Delete(mysql.DB().GetDbW().WithContext(ctx)); err != nil {
		return 0, err
	}

	// 建立新关系
	list := make([]*ums_admin_role_relation.UmsAdminRoleRelation, 0, count)
	for _, roleId := range roleIds {
		roleRelation := ums_admin_role_relation.NewModel()
		roleRelation.AdminId = adminId
		roleRelation.RoleId = roleId
		list = append(list, roleRelation)
	}
	roleRelationDao := new(dao.UmsAdminRoleRelationDao)
	err := roleRelationDao.InsertList(mysql.DB().GetDbW().WithContext(ctx), list)
	if err != nil {
		return 0, err
	}
	s.cacheService.DelResourceList(ctx, adminId)
	return count, nil
}
