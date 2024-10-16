package sms_home_advertise

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/pagehelper"
)

var _ Service = (*service)(nil)

// 首页广告管理Service
type Service interface {
	i()

	/**
	 * 添加广告
	 */
	Create(ctx context.Context, param dto.SmsHomeAdvertise) (int64, error)

	/**
	 * 批量删除广告
	 */
	Delete(ctx context.Context, ids []int64) (int64, error)

	/**
	 * 修改上、下线状态
	 */
	UpdateStatus(ctx context.Context, id int64, status int32) (int64, error)

	/**
	 * 获取广告详情
	 */
	GetItem(ctx context.Context, id int64) (*dto.SmsHomeAdvertise, error)

	/**
	 * 更新广告
	 */
	Update(ctx context.Context, id int64, param dto.SmsHomeAdvertise) (int64, error)

	/**
	 * 分页查询广告
	 */
	List(ctx context.Context, name string, adType int32, endTime string, pageSize, pageNum int) (
		*pagehelper.ListData[dto.SmsHomeAdvertise], error)
}
