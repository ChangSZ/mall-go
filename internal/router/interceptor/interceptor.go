package interceptor

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/internal/services/authorized"

	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	CheckToken(ctx core.Context) (info proposal.UmsUserInfo, err core.BusinessError)

	// CheckRBAC 验证 RBAC 权限是否合法
	CheckRBAC() core.HandlerFunc

	// CheckSignature 验证签名是否合法，对用签名算法 pkg/signature
	CheckSignature() core.HandlerFunc

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger            *zap.Logger
	db                mysql.Repo
	authorizedService authorized.Service
	adminService      admin.Service
}

func New(logger *zap.Logger, db mysql.Repo) Interceptor {
	return &interceptor{
		logger:            logger,
		db:                db,
		authorizedService: authorized.New(db),
		adminService:      admin.New(db),
	}
}

func (i *interceptor) i() {}
