package ums_member_level

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 会员等级管理Service
type Service interface {
	i()

	/**
	 * 获取所有会员等级
	 * @param defaultStatus 是否为默认会员
	 */
	List(ctx context.Context, defaultStatus int32) ([]dto.UmsMemberLevel, error)
}
