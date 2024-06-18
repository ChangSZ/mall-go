package alipay

import (
	"context"
	"net/url"

	"github.com/ChangSZ/mall-go/internal/dto"

	"github.com/smartwalle/alipay/v3"
)

var _ Service = (*service)(nil)

// 支付宝支付Service
type Service interface {
	i()

	/**
	 * 根据提交参数生成电脑支付页面
	 */
	Pay(ctx context.Context, aliPayParam dto.AliPayParam) (string, error)

	/**
	 * 支付宝异步回调处理
	 */
	Notify(ctx context.Context, param url.Values) (string, error)

	/**
	 * 查询支付宝交易状态
	 * @param outTradeNo 商户订单编号
	 * @param tradeNo 支付宝交易编号
	 * @return 支付宝交易状态
	 */
	Query(ctx context.Context, outTradeNo, tradeNo string) (alipay.TradeStatus, error)

	/**
	 * 根据提交参数生成手机支付页面
	 */
	PhoneWebPay(ctx context.Context, aliPayParam dto.AliPayParam) (string, error)
}
