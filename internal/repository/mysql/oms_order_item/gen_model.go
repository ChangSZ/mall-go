package oms_order_item

// OmsOrderItem 订单中所包含的商品
//
//go:generate gormgen -structs OmsOrderItem -input .
type OmsOrderItem struct {
	Id                int64   //
	OrderId           int64   // 订单id
	OrderSn           string  // 订单编号
	ProductId         int64   //
	ProductPic        string  //
	ProductName       string  //
	ProductBrand      string  //
	ProductSn         string  //
	ProductPrice      float64 // 销售价格
	ProductQuantity   int32   // 购买数量
	ProductSkuId      int64   // 商品sku编号
	ProductSkuCode    string  // 商品sku条码
	ProductCategoryId int64   // 商品分类id
	PromotionName     string  // 商品促销名称
	PromotionAmount   float64 // 商品促销分解金额
	CouponAmount      float64 // 优惠券优惠分解金额
	IntegrationAmount float64 // 积分优惠分解金额
	RealAmount        float64 // 该商品经过优惠后的分解金额
	GiftIntegration   int32   //
	GiftGrowth        int32   //
	ProductAttr       string  // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
