package pms_product

import "time"

// PmsProduct 商品信息
//
//go:generate gormgen -structs PmsProduct -input .
type PmsProduct struct {
	Id                         int64     //
	BrandId                    int64     //
	ProductCategoryId          int64     //
	FeightTemplateId           int64     //
	ProductAttributeCategoryId int64     //
	Name                       string    //
	Pic                        string    //
	ProductSn                  string    // 货号
	DeleteStatus               int32     // 删除状态：0->未删除；1->已删除
	PublishStatus              int32     // 上架状态：0->下架；1->上架
	NewStatus                  int32     // 新品状态:0->不是新品；1->新品
	RecommandStatus            int32     // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int32     // 审核状态：0->未审核；1->审核通过
	Sort                       int32     // 排序
	Sale                       int32     // 销量
	Price                      float64   //
	PromotionPrice             float64   // 促销价格
	GiftGrowth                 int32     // 赠送的成长值
	GiftPoint                  int32     // 赠送的积分
	UsePointLimit              int32     // 限制使用的积分数
	SubTitle                   string    // 副标题
	Description                string    // 商品描述
	OriginalPrice              float64   // 市场价
	Stock                      int32     // 库存
	LowStock                   int32     // 库存预警值
	Unit                       string    // 单位
	Weight                     float64   // 商品重量，默认为克
	PreviewStatus              int32     // 是否为预告商品：0->不是；1->是
	ServiceIds                 string    // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string    //
	Note                       string    //
	AlbumPics                  string    // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string    //
	DetailDesc                 string    //
	DetailHtml                 string    // 产品详情网页内容
	DetailMobileHtml           string    // 移动端网页详情
	PromotionStartTime         time.Time `gorm:"time"` // 促销开始时间
	PromotionEndTime           time.Time `gorm:"time"` // 促销结束时间
	PromotionPerLimit          int32     // 活动限购数量
	PromotionType              int32     // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string    // 品牌名称
	ProductCategoryName        string    // 商品分类名称
}
