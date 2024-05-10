package dto

type OmsOrderItem struct {
	Id                int64   `json:"id"`                //
	OrderId           int64   `json:"orderId"`           // 订单id
	OrderSn           string  `json:"orderSn"`           // 订单编号
	ProductId         int64   `json:"productId"`         //
	ProductPic        string  `json:"productPic"`        //
	ProductName       string  `json:"productName"`       //
	ProductBrand      string  `json:"productBrand"`      //
	ProductSn         string  `json:"productSn"`         //
	ProductPrice      float64 `json:"productPrice"`      // 销售价格
	ProductQuantity   int32   `json:"productQuantity"`   // 购买数量
	ProductSkuId      int64   `json:"productSkuId"`      // 商品sku编号
	ProductSkuCode    string  `json:"productSkuCode"`    // 商品sku条码
	ProductCategoryId int64   `json:"productCategoryId"` // 商品分类id
	PromotionName     string  `json:"promotionName"`     // 商品促销名称
	PromotionAmount   float64 `json:"promotionAmount"`   // 商品促销分解金额
	CouponAmount      float64 `json:"couponAmount"`      // 优惠券优惠分解金额
	IntegrationAmount float64 `json:"integrationAmount"` // 积分优惠分解金额
	RealAmount        float64 `json:"realAmount"`        // 该商品经过优惠后的分解金额
	GiftIntegration   int32   `json:"giftIntegration"`   //
	GiftGrowth        int32   `json:"giftGrowth"`        //
	ProductAttr       string  `json:"productAttr"`       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
