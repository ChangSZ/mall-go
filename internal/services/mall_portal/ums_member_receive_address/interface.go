package ums_member_receive_address

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 用户地址管理Service
type Service interface {
	i()

	/**
	 * 添加收货地址
	 */
	Add(ctx context.Context, param dto.UmsMemberReceiveAddress) (int64, error)

	/**
	 * 删除收货地址
	 * @param id 地址表的id
	 */
	Delete(ctx context.Context, id int64) (int64, error)

	/**
	 * 修改收货地址
	 * @param id 地址表的id
	 * @param address 修改的收货地址信息
	 */
	Update(ctx context.Context, id int64, param dto.UmsMemberReceiveAddress) (int64, error)

	/**
	 * 返回当前用户的收货地址
	 */
	List(ctx context.Context) ([]dto.UmsMemberReceiveAddress, error)

	/**
	 * 获取地址详情
	 * @param id 地址id
	 */
	GetItem(ctx context.Context, id int64) (*dto.UmsMemberReceiveAddress, error)
}
