package cms_prefrence_area

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 优选专区管理Service
type Service interface {
	i()

	/**
	 * 获取所有优选专区
	 */
	ListAll(ctx context.Context) ([]dto.CmsPrefrenceArea, error)
}
