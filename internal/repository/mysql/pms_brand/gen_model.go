package pms_brand

// PmsBrand 品牌表
//
//go:generate gormgen -structs PmsBrand -input .
type PmsBrand struct {
	Id                  int64  //
	Name                string //
	FirstLetter         string // 首字母
	Sort                int32  //
	FactoryStatus       int32  // 是否为品牌制造商：0->不是；1->是
	ShowStatus          int32  //
	ProductCount        int32  // 产品数量
	ProductCommentCount int32  // 产品评论数量
	Logo                string // 品牌logo
	BigPic              string // 专区大图
	BrandStory          string // 品牌故事
}
