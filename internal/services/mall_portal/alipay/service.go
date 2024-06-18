package alipay

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ChangSZ/mall-go/configs"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/services/mall_portal/oms_portal_order"
	"github.com/ChangSZ/mall-go/pkg/env"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/smartwalle/alipay/v3"
)

var (
	GatewayUrl      = configs.Get().Alipay.GatewayUrl
	AppId           = configs.Get().Alipay.AppId
	AlipayPublicKey = configs.Get().Alipay.AlipayPublicKey
	AppPrivateKey   = configs.Get().Alipay.AppPrivateKey
	ApiAESKey       = configs.Get().Alipay.ApiAESKey
	ReturnUrl       = configs.Get().Alipay.ReturnUrl
	NotifyUrl       = configs.Get().Alipay.NotifyUrl
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) GetClient() (*alipay.Client, error) {
	opt := alipay.WithNewSandboxGateway()
	if GatewayUrl == "https://openapi.alipaydev.com/gateway.do" {
		opt = alipay.WithPastSandboxGateway()
	}
	client, err := alipay.New(AppId, AppPrivateKey, env.Active().IsPro(), opt)
	if err != nil {
		return nil, fmt.Errorf("初始化支付宝失败: %v", err)
	}

	if err := client.LoadAliPayPublicKey(AlipayPublicKey); err != nil {
		return nil, fmt.Errorf("加载支付宝公钥时发生错误: %v", err)
	}

	if err := client.SetEncryptKey(ApiAESKey); err != nil {
		return nil, fmt.Errorf("加载内容加密密钥发生错误: %v", err)
	}
	return client, nil
}

func (s *service) Pay(ctx context.Context, aliPayParam dto.AliPayParam) (string, error) {
	client, err := s.GetClient()
	if err != nil {
		return "", err
	}

	var p = alipay.TradePagePay{}
	// 异步接收地址，公网可访问
	// p.NotifyURL = NotifyUrl
	// 同步跳转地址
	p.ReturnURL = ReturnUrl
	// 订单标题，不可使用特殊符号
	p.Subject = aliPayParam.Subject
	// 商户订单号，商家自定义，保持唯一性
	p.OutTradeNo = aliPayParam.OutTradeNo
	// 支付金额，最小值0.01元
	p.TotalAmount = fmt.Sprintf("%.2f", aliPayParam.TotalAmount)
	// 电脑网站支付场景固定传值FAST_INSTANT_TRADE_PAY
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (s *service) Notify(ctx context.Context, param url.Values) (string, error) {
	result := "failure"
	client, err := s.GetClient()
	if err != nil {
		return result, err
	}

	// DecodeNotification 内部已调用 VerifySign 方法验证签名
	noti, err := client.DecodeNotification(param)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		return result, fmt.Errorf("支付回调签名校验失败")
	}

	if noti.TradeStatus == alipay.TradeStatusSuccess {
		if err := oms_portal_order.New().PaySuccessByOrderSn(ctx, noti.OutTradeNo, 1); err != nil {
			log.WithTrace(ctx).Error(err)
			return result, err
		}
	} else {
		return result, fmt.Errorf("订单未支付成功, trade_status: %v", noti.TradeStatus)
	}
	// 如果通知消息没有问题，我们需要确认收到通知消息，不然支付宝后续会继续推送相同的消息
	result = "success"
	return result, nil
}

func (s *service) Query(ctx context.Context, outTradeNo, tradeNo string) (alipay.TradeStatus, error) {
	client, err := s.GetClient()
	if err != nil {
		return "", err
	}

	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	p.TradeNo = tradeNo
	// 交易结算信息: trade_settle_info
	p.QueryOptions = []string{"TRADE_SETTLE_INFO"}
	rsp, err := client.TradeQuery(context.Background(), p)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		return "", fmt.Errorf("查询支付宝账单异常")
	}

	if rsp.IsFailure() {
		return "", fmt.Errorf("查询支付宝账单失败")
	}

	log.WithTrace(ctx).Info("查询支付宝账单成功")
	if rsp.TradeStatus == alipay.TradeStatusSuccess {
		if err := oms_portal_order.New().PaySuccessByOrderSn(ctx, outTradeNo, 1); err != nil {
			log.WithTrace(ctx).Error(err)
			return "", err
		}
	}

	/*
		交易状态：
			WAIT_BUYER_PAY: 交易创建，等待买家付款
			TRADE_CLOSED: 未付款交易超时关闭，或支付完成后全额退款
			TRADE_SUCCESS: 交易支付成功
			TRADE_FINISHED: 交易结束，不可退款
	*/
	return rsp.TradeStatus, nil
}

func (s *service) PhoneWebPay(ctx context.Context, aliPayParam dto.AliPayParam) (string, error) {
	client, err := s.GetClient()
	if err != nil {
		return "", err
	}

	var p = alipay.TradeWapPay{}
	// 异步接收地址，公网可访问
	// p.NotifyURL = NotifyUrl
	// 同步跳转地址
	p.ReturnURL = ReturnUrl
	// 订单标题，不可使用特殊符号
	p.Subject = aliPayParam.Subject
	// 商户订单号，商家自定义，保持唯一性
	p.OutTradeNo = aliPayParam.OutTradeNo
	// 支付金额，最小值0.01元
	p.TotalAmount = fmt.Sprintf("%.2f", aliPayParam.TotalAmount)
	// 手机网站支付默认传值QUICK_WAP_WAY
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
