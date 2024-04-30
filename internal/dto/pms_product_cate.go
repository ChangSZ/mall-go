package dto

type PmsProductCategoryParam struct {
	ParentId               int64   `json:"parentId" binding:"required"`              // 父分类的编号
	Name                   string  `json:"name" binding:"required"`                  // 商品分类名称
	ProductUnit            string  `json:"productUnit"`                              // 分类单位
	NavStatus              int     `json:"navStatus" binding:"omitempty,oneof=0 1"`  // 是否在导航栏显示
	ShowStatus             int     `json:"showStatus" binding:"omitempty,oneof=0 1"` // 是否进行显示
	Sort                   int     `json:"sort" binding:"gte=0"`                     // 排序
	Icon                   string  `json:"icon"`                                     // 图标
	Keywords               string  `json:"keywords"`                                 // 关键字
	Description            string  `json:"description"`                              // 描述
	ProductAttributeIdList []int64 `json:"productAttributeIdList"`                   // 产品相关筛选属性集合
}

type PmsProductCategory struct {
	Id           int64  `json:"id"`
	ParentId     int64  `json:"parentId"`
	Name         string `json:"name"`
	Level        int32  `json:"level"`
	ProductCount int32  `json:"productCount"`
	ProductUnit  string `json:"productUnit"`
	NavStatus    int32  `json:"navStatus"`
	ShowStatus   int32  `json:"showStatus"`
	Sort         int32  `json:"sort"`
	Icon         string `json:"icon"`
	Keywords     string `json:"keywords"`
	Description  string `json:"description"`
}

// PmsProductCategoryWithChildrenItem 包含子级分类的商品分类
type PmsProductCategoryWithChildrenItem struct {
	PmsProductCategory `json:",inline"`
	Children           []PmsProductCategory `json:"children" gorm:"foreignKey:ParentId"` // 子级分类
}

type PmsProductCateUri struct {
	ParentId int64 `uri:"parentId" binding:"required"`
}
