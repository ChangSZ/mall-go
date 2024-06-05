package oms_cart_item

import "time"

// OmsCartItem 购物车表
//
//go:generate gormgen -structs OmsCartItem -input .
type OmsCartItem struct {
	Id                int64     //
	ProductId         int64     //
	ProductSkuId      int64     //
	MemberId          int64     //
	Quantity          int32     // 购买数量
	Price             float64   // 添加到购物车的价格
	ProductPic        string    // 商品主图
	ProductName       string    // 商品名称
	ProductSubTitle   string    // 商品副标题（卖点）
	ProductSkuCode    string    // 商品sku条码
	MemberNickname    string    // 会员昵称
	CreateDate        time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 创建时间
	ModifyDate        time.Time `gorm:"type:time;default:'1000-01-01 00:00:00'"` // 修改时间
	DeleteStatus      int32     // 是否删除
	ProductCategoryId int64     // 商品分类
	ProductBrand      string    //
	ProductSn         string    //
	ProductAttr       string    // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}
