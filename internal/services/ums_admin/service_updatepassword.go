package ums_admin

import (
	"github.com/ChangSZ/mall-go/internal/pkg/core"
)

type UpdateAdminPasswordParam struct {
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

func (s *service) UpdatePassword(ctx core.Context, updatePasswordParam *UpdateAdminPasswordParam) (int64, error) {
	return 0, nil
}
