package ums_member

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member_level"
	"github.com/ChangSZ/mall-go/pkg/jwt"
	"github.com/ChangSZ/mall-go/pkg/password"
)

var (
	jwtConfig    = configs.Get().Jwt
	jwtTokenUtil = jwt.NewJwtTokenUtil(jwtConfig.Secret, jwtConfig.Expiration, jwtConfig.TokenHead)
)

type service struct {
	cacheService *UmsMemberCacheService
}

func New() Service {
	return &service{NewCacheService()}
}

func (s *service) i() {}

func (s *service) GetByUsername(ctx context.Context, username string) (*ums_member.UmsMember, error) {
	member := s.cacheService.GetMember(ctx, username)
	if member != nil {
		return member, nil
	}
	qb := ums_member.NewQueryBuilder()
	qb = qb.WhereUsername(mysql.EqualPredicate, username)
	memberList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if len(memberList) == 0 {
		return nil, fmt.Errorf("未找到该用户: %s", username)
	}
	member = memberList[0]
	if len(memberList) > 1 {
		return member, fmt.Errorf("存在多个(%d)用户, 请联系管理员", len(memberList))
	}
	return member, nil
}

func (s *service) GetById(ctx context.Context, id int64) (*ums_member.UmsMember, error) {
	qb := ums_member.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	return qb.First(mysql.DB().GetDbR())
}

func (s *service) Register(ctx context.Context, username, passwd, telephone, authCode string) error {
	// 验证验证码
	if !s.verifyAuthCode(ctx, telephone, authCode) {
		return fmt.Errorf("验证码错误")
	}
	// 查询是否已有该用户
	umsMembers, err := new(dao.UmsMemberDao).GetMembers(ctx, mysql.DB().GetDbR().WithContext(ctx), username, telephone)
	if err != nil {
		return err
	}
	if len(umsMembers) > 0 {
		return fmt.Errorf("该用户已经存在")
	}

	encodePassword, err := password.Encoder.Encode(passwd)
	if err != nil {
		return fmt.Errorf("failed to encode password: %v", err)
	}
	data := ums_member.NewModel()
	data.Username = username
	data.Phone = telephone
	data.Password = encodePassword
	data.Status = 1

	// 获取默认会员等级并设置
	qb := ums_member_level.NewQueryBuilder()
	qb = qb.WhereDefaultStatus(mysql.EqualPredicate, 1)
	memberLevelList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if len(memberLevelList) > 0 {
		data.MemberLevelId = memberLevelList[0].Id
	}
	_, err = data.Create(mysql.DB().GetDbW().WithContext(ctx))
	data.Password = ""
	return err
}

func (s *service) GenerateAuthCode(ctx context.Context, telephone string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	authCode := ""
	for i := 0; i < 6; i++ {
		authCode += fmt.Sprintf("%d", rand.Intn(10))
	}
	s.cacheService.SetAuthCode(ctx, telephone, authCode)
	return authCode
}

func (s *service) UpdatePassword(ctx context.Context, telephone, passwd, authCode string) error {
	qb := ums_member.NewQueryBuilder()
	qb = qb.WherePhone(mysql.EqualPredicate, telephone)
	memberList, err := qb.QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return err
	}
	if len(memberList) == 0 {
		return fmt.Errorf("该账号不存在")
	}
	if len(memberList) > 1 {
		return fmt.Errorf("异常: 存在多个账号, 请联系管理员")
	}

	// 验证验证码
	if !s.verifyAuthCode(ctx, telephone, authCode) {
		return fmt.Errorf("验证码错误")
	}
	encodePassword, err := password.Encoder.Encode(passwd)
	if err != nil {
		return fmt.Errorf("failed to encode password: %v", err)
	}

	id := memberList[0].Id
	data := map[string]interface{}{"password": encodePassword}
	{
		qb := ums_member.NewQueryBuilder()
		qb = qb.WhereId(mysql.EqualPredicate, id)
		_, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
		if err != nil {
			return err
		}
		s.cacheService.DelMember(ctx, id)
	}
	return nil
}

func (s *service) UpdateIntegration(ctx context.Context, id int64, integration int32) error {
	data := map[string]interface{}{
		"integration": integration,
	}
	qb := ums_member.NewQueryBuilder()
	qb = qb.WhereId(mysql.EqualPredicate, id)
	_, err := qb.Updates(mysql.DB().GetDbW().WithContext(ctx), data)
	if err != nil {
		return err
	}
	s.cacheService.DelMember(ctx, id)
	return nil
}

func (s *service) LoadUserByUsername(ctx context.Context, username string) (*MemberUserDetails, error) {
	// 获取用户信息
	member, err := s.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if member != nil {
		return &MemberUserDetails{member}, nil
	}
	return nil, fmt.Errorf("用户名或密码错误")
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

func (s *service) verifyAuthCode(ctx context.Context, telephone, authCode string) bool {
	if authCode == "" {
		return false
	}
	return authCode == s.cacheService.GetAuthCode(ctx, telephone)
}
