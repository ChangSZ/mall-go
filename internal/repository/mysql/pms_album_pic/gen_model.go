package pms_album_pic

// PmsAlbumPic 画册图片表
//
//go:generate gormgen -structs PmsAlbumPic -input .
type PmsAlbumPic struct {
	Id      int64  //
	AlbumId int64  //
	Pic     string //
}
