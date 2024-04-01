package ums_admin

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_admin"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_resource"
)

type AdminUserDetails struct {
	UmsAdmin     ums_admin.UmsAdmin
	ResourceList []ums_resource.UmsResource
}

func (au *AdminUserDetails) GetAuthorities() []string {
	authorities := make([]string, len(au.ResourceList))
	for i, resource := range au.ResourceList {
		authorities[i] = fmt.Sprintf("%d:%s", resource.Id, resource.Name)
	}
	return authorities
}

func (au *AdminUserDetails) GetPassword() string {
	return au.UmsAdmin.Password
}

func (au *AdminUserDetails) GetUsername() string {
	return au.UmsAdmin.Username
}

func (au *AdminUserDetails) IsAccountNonExpired() bool {
	return true
}

func (au *AdminUserDetails) IsAccountNonLocked() bool {
	return true
}

func (au *AdminUserDetails) IsCredentialsNonExpired() bool {
	return true
}

func (au *AdminUserDetails) IsEnabled() bool {
	return au.UmsAdmin.Status == 1
}
