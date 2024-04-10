package interceptor

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
	"github.com/ChangSZ/mall-go/internal/proposal"
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	CheckToken(ctx core.Context) (info proposal.UmsUserInfo, err core.BusinessError)

	// CheckRBAC 验证 RBAC 权限是否合法
	CheckRBAC(*gin.Context)

	// CheckSignature 验证签名是否合法，对用签名算法 pkg/signature
	CheckSignature(*gin.Context)

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger            *zap.Logger
	authorizedService authorized.Service
	adminService      admin.Service
}

func New(logger *zap.Logger) Interceptor {
	return &interceptor{
		logger:            logger,
		authorizedService: authorized.New(),
		adminService:      admin.New(),
	}
}

func (i *interceptor) i() {}
