package cms_subject

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 商品专题管理Service
type Service interface {
	i()

	/**
	 * 查询所有专题
	 */
	ListAll(ctx context.Context) ([]dto.CmsSubject, error)

	/**
	 * 分页查询专题
	 */
	List(ctx context.Context, keyword string, pageSize, pageNum int) (*pagehelper.ListData[dto.CmsSubject], error)
}
