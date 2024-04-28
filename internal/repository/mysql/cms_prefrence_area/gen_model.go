package cms_prefrence_area

// CmsPrefrenceArea 优选专区
//
//go:generate gormgen -structs CmsPrefrenceArea -input .
type CmsPrefrenceArea struct {
	Id         int64  //
	Name       string //
	SubTitle   string //
	Pic        []byte // 展示图片
	Sort       int32  //
	ShowStatus int32  //
}
