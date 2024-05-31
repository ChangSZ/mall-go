package dto

import "time"

type OmsCartItem struct {
	Id                int64     `json:"id"`                //
	ProductId         int64     `json:"productId"`         //
	ProductSkuId      int64     `json:"productSkuId"`      //
	MemberId          int64     `json:"memberId"`          //
	Quantity          int32     `json:"quantity"`          // 购买数量
	Price             float64   `json:"price"`             // 添加到购物车的价格
	ProductPic        string    `json:"productPic"`        // 商品主图
	ProductName       string    `json:"productName"`       // 商品名称
	ProductSubTitle   string    `json:"productSubTitle"`   // 商品副标题（卖点）
	ProductSkuCode    string    `json:"productSkuCode"`    // 商品sku条码
	MemberNickname    string    `json:"memberNickname"`    // 会员昵称
	CreateDate        time.Time `json:"createDate"`        // 创建时间
	ModifyDate        time.Time `json:"modifyDate"`        // 修改时间
	DeleteStatus      int32     `json:"deleteStatus"`      // 是否删除
	ProductCategoryId int64     `json:"productCategoryId"` // 商品分类
	ProductBrand      string    `json:"productBrand"`      //
	ProductSn         string    `json:"productSn"`         //
	ProductAttr       string    `json:"productAttr"`       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
