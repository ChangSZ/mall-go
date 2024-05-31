package ums_member

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	/**
	 * 根据用户名获取会员
	 */
	GetByUsername(ctx context.Context, username string) (*ums_member.UmsMember, error)

	/**
	 * 根据会员编号获取会员
	 */
	GetById(ctx context.Context, id int64) (*ums_member.UmsMember, error)

	/**
	 * 用户注册
	 */
	Register(ctx context.Context, username, passwd, telephone, authCode string) error

	/**
	 * 生成验证码
	 */
	GenerateAuthCode(ctx context.Context, telephone string) string

	/**
	 * 修改密码
	 */
	UpdatePassword(ctx context.Context, telephone, passwd, authCode string) error

	/**
	 * 根据会员id修改会员积分
	 */
	UpdateIntegration(ctx context.Context, id int64, integration int32) error

	/**
	 * 获取当前登录会员
	 */
	GetCurrentMember(ctx context.Context) (*dto.UmsMember, error)

	/**
	 * 获取用户信息
	 */
	LoadUserByUsername(ctx context.Context, username string) (*MemberUserDetails, error)

	/**
	 * 登录后获取token
	 */
	Login(ctx context.Context, username, passwd string) (string, error)

	/**
	 * 刷新token
	 */
	RefreshToken(ctx context.Context, oldToken string) (string, error)
}

// 会员信息缓存业务类Service
type UmsMemberCacheServiceI interface {

	/**
	 * 删除会员用户缓存
	 */
	DelMember(ctx context.Context, memberId int64)

	/**
	 * 获取会员用户缓存
	 */
	GetMember(ctx context.Context, username string) *ums_member.UmsMember

	/**
	 * 设置会员用户缓存
	 */
	SetMember(ctx context.Context, member *ums_member.UmsMember)

	/**
	 * 设置验证码
	 */
	SetAuthCode(ctx context.Context, telephone, authCode string)

	/**
	 * 获取验证码
	 */
	GetAuthCode(ctx context.Context, telephone string) string
}
