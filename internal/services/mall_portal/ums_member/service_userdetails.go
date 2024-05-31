package ums_member

import (
	"github.com/ChangSZ/mall-go/internal/repository/mysql/ums_member"
)

type MemberUserDetails struct {
	UmsMember *ums_member.UmsMember
}

func (au *MemberUserDetails) GetPassword() string {
	return au.UmsMember.Password
}

func (au *MemberUserDetails) GetUsername() string {
	return au.UmsMember.Username
}

func (au *MemberUserDetails) IsAccountNonExpired() bool {
	return true
}

func (au *MemberUserDetails) IsAccountNonLocked() bool {
	return true
}

func (au *MemberUserDetails) IsCredentialsNonExpired() bool {
	return true
}

func (au *MemberUserDetails) IsEnabled() bool {
	return au.UmsMember.Status == 1
}
