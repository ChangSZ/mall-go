package pms_product_category

// PmsProductCategory 产品分类
//
//go:generate gormgen -structs PmsProductCategory -input .
type PmsProductCategory struct {
	Id           int64  //
	ParentId     int64  // 上机分类的编号：0表示一级分类
	Name         string //
	Level        int32  // 分类级别：0->1级；1->2级
	ProductCount int32  //
	ProductUnit  string //
	NavStatus    int32  // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int32  // 显示状态：0->不显示；1->显示
	Sort         int32  //
	Icon         string // 图标
	Keywords     string //
	Description  string // 描述
}
