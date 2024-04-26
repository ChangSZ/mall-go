package pms_brand

type PmsBrandParam struct {
	Name                string `json:"name" binding:"required"`
	FirstLetter         string `json:"firstLetter"`
	Sort                int32  `json:"sort" binding:"omitempty,gte=0"`
	FactoryStatus       int32  `json:"factoryStatus" binding:"omitempty,oneof=0 1"`
	ShowStatus          int32  `json:"showStatus"  binding:"omitempty,oneof=0 1"`
	ProductCount        int32  `json:"productCount"`
	ProductCommentCount int32  `json:"productCommentCount"`
	Logo                string `json:"logo" binding:"required"`
	BigPic              string `json:"bigPic"`
	BrandStory          string `json:"brandStory"`
}

type PmsBrand struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	FirstLetter         string `json:"firstLetter"`
	Sort                int32  `json:"sort"`
	FactoryStatus       int32  `json:"factoryStatus"`
	ShowStatus          int32  `json:"showStatus"`
	ProductCount        int32  `json:"productCount"`
	ProductCommentCount int32  `json:"productCommentCount"`
	Logo                string `json:"logo"`
	BigPic              string `json:"bigPic"`
	BrandStory          string `json:"brandStory"`
}

type PmsBrandUri struct {
	Id int64 `uri:"id" binding:"required"`
}
