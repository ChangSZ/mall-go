package dto

import "time"

type PmsProductResult struct {
	PmsProductParam `json:",inline"`
	CateParentId    int64 `json:"cateParentId"` // 商品所选分类的父id
}

type PmsProductParam struct {
	PmsProduct                       `json:",inline"`
	ProductLadderList                []PmsProductLadder                `json:"productLadderList" gorm:"foreignKey:ProductId"`                // 商品阶梯价格设置
	ProductFullReductionList         []PmsProductFullReduction         `json:"productFullReductionList" gorm:"foreignKey:ProductId"`         // 商品满减价格设置
	MemberPriceList                  []PmsMemberPrice                  `json:"memberPriceList" gorm:"foreignKey:ProductId"`                  // 商品会员价格设置
	SkuStockList                     []PmsSkuStock                     `json:"skuStockList" gorm:"foreignKey:ProductId"`                     // 商品的sku库存信息
	ProductAttributeValueList        []PmsProductAttributeValue        `json:"productAttributeValueList" gorm:"foreignKey:ProductId"`        // 商品参数及自定义规格属性
	SubjectProductRelationList       []CmsSubjectProductRelation       `json:"subjectProductRelationList" gorm:"foreignKey:ProductId"`       // 专题和商品关系
	PrefrenceAreaProductRelationList []CmsPrefrenceAreaProductRelation `json:"prefrenceAreaProductRelationList" gorm:"foreignKey:ProductId"` // 优选专区和商品的关系
}

type PmsProductQueryParam struct {
	PublishStatus     int32  `form:"publishStatus" json:"publishStatus,omitempty"`         // 上架状态
	VerifyStatus      int32  `form:"verifyStatus" json:"verifyStatus,omitempty"`           // 审核状态
	Keyword           string `form:"keyword" json:"keyword,omitempty"`                     // 商品名称模糊关键字
	ProductSn         string `form:"productSn" json:"productSn,omitempty"`                 // 商品货号
	ProductCategoryId int64  `form:"productCategoryId" json:"productCategoryId,omitempty"` // 商品分类编号
	BrandId           int64  `form:"brandId" json:"brandId,omitempty"`                     // 商品品牌编号
}

type PmsProduct struct {
	Id                         int64     `json:"id"`                         //
	BrandId                    int64     `json:"brandId"`                    //
	ProductCategoryId          int64     `json:"productCategoryId"`          //
	FeightTemplateId           int64     `json:"feightTemplateId"`           //
	ProductAttributeCategoryId int64     `json:"productAttributeCategoryId"` //
	Name                       string    `json:"name"`                       //
	Pic                        string    `json:"pic"`                        //
	ProductSn                  string    `json:"productSn"`                  // 货号
	DeleteStatus               int32     `json:"deleteStatus"`               // 删除状态：0->未删除；1->已删除
	PublishStatus              int32     `json:"publishStatus"`              // 上架状态：0->下架；1->上架
	NewStatus                  int32     `json:"newStatus"`                  // 新品状态:0->不是新品；1->新品
	RecommandStatus            int32     `json:"recommandStatus"`            // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int32     `json:"verifyStatus"`               // 审核状态：0->未审核；1->审核通过
	Sort                       int32     `json:"sort"`                       // 排序
	Sale                       int32     `json:"sale"`                       // 销量
	Price                      float64   `json:"price"`                      //
	PromotionPrice             float64   `json:"promotionPrice"`             // 促销价格
	GiftGrowth                 int32     `json:"giftGrowth"`                 // 赠送的成长值
	GiftPoint                  int32     `json:"giftPoint"`                  // 赠送的积分
	UsePointLimit              int32     `json:"usePointLimit"`              // 限制使用的积分数
	SubTitle                   string    `json:"subTitle"`                   // 副标题
	Description                string    `json:"description"`                // 商品描述
	OriginalPrice              float64   `json:"originalPrice"`              // 市场价
	Stock                      int32     `json:"stock"`                      // 库存
	LowStock                   int32     `json:"lowStock"`                   // 库存预警值
	Unit                       string    `json:"unit"`                       // 单位
	Weight                     float64   `json:"weight"`                     // 商品重量，默认为克
	PreviewStatus              int32     `json:"previewStatus"`              // 是否为预告商品：0->不是；1->是
	ServiceIds                 string    `json:"serviceIds"`                 // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string    `json:"keywords"`                   //
	Note                       string    `json:"note"`                       //
	AlbumPics                  string    `json:"albumPics"`                  // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string    `json:"detailTitle"`                //
	DetailDesc                 string    `json:"detailDesc"`                 //
	DetailHtml                 string    `json:"detailHtml"`                 // 产品详情网页内容
	DetailMobileHtml           string    `json:"detailMobileHtml"`           // 移动端网页详情
	PromotionStartTime         time.Time `json:"promotionStartTime"`         // 促销开始时间
	PromotionEndTime           time.Time `json:"promotionEndTime"`           // 促销结束时间
	PromotionPerLimit          int32     `json:"promotionPerLimit"`          // 活动限购数量
	PromotionType              int32     `json:"promotionType"`              // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string    `json:"brandName"`                  // 品牌名称
	ProductCategoryName        string    `json:"productCategoryName"`        // 商品分类名称
}

type PmsProductLadder struct {
	Id        int64   `json:"id"`        //
	ProductId int64   `json:"productId"` //
	Count     int32   `json:"count"`     // 满足的商品数量
	Discount  float64 `json:"discount"`  // 折扣
	Price     float64 `json:"price"`     // 折后价格
}

type PmsProductFullReduction struct {
	Id          int64   `json:"id"`          //
	ProductId   int64   `json:"productId"`   //
	FullPrice   float64 `json:"fullPrice"`   //
	ReducePrice float64 `json:"reducePrice"` //
}

type PmsMemberPrice struct {
	Id              int64   `json:"id"`              //
	ProductId       int64   `json:"productId"`       //
	MemberLevelId   int64   `json:"memberLevelId"`   //
	MemberPrice     float64 `json:"memberPrice"`     // 会员价格
	MemberLevelName string  `json:"memberLevelName"` //
}

// PmsPortalProductDetail 前台商品详情
type PmsPortalProductDetail struct {
	Product                   PmsProduct                 `json:"product"`                   // 商品信息
	Brand                     PmsBrand                   `json:"brand"`                     // 商品品牌
	ProductAttributeList      []PmsProductAttribute      `json:"productAttributeList"`      // 商品属性与参数
	ProductAttributeValueList []PmsProductAttributeValue `json:"productAttributeValueList"` // 手动录入的商品属性与参数值
	SkuStockList              []PmsSkuStock              `json:"skuStockList"`              // 商品的sku库存信息
	ProductLadderList         []PmsProductLadder         `json:"productLadderList"`         // 商品阶梯价格设置
	ProductFullReductionList  []PmsProductFullReduction  `json:"productFullReductionList"`  // 商品满减价格设置
	CouponList                []SmsCoupon                `json:"couponList"`                // 商品可用优惠券
}
