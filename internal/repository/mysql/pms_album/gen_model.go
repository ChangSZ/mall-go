package pms_album

// PmsAlbum 相册表
//
//go:generate gormgen -structs PmsAlbum -input .
type PmsAlbum struct {
	Id          int64  //
	Name        string //
	CoverPic    string //
	PicCount    int32  //
	Sort        int32  //
	Description string //
}
