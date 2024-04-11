package interceptor

import (
	"github.com/ChangSZ/mall-go/internal/services/admin"
	"github.com/ChangSZ/mall-go/internal/services/authorized"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	authorizedService authorized.Service
	adminService      admin.Service
}

func New() Interceptor {
	return &interceptor{
		authorizedService: authorized.New(),
		adminService:      admin.New(),
	}
}

func (i *interceptor) i() {}
