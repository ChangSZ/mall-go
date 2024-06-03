package dto

type UmsMemberReceiveAddress struct {
	Id            int64  `json:"id"`            //
	MemberId      int64  `json:"memberId"`      //
	Name          string `json:"name"`          // 收货人名称
	PhoneNumber   string `json:"phoneNumber"`   //
	DefaultStatus int32  `json:"defaultStatus"` // 是否为默认
	PostCode      string `json:"postCode"`      // 邮政编码
	Province      string `json:"province"`      // 省份/直辖市
	City          string `json:"city"`          // 城市
	Region        string `json:"region"`        // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}
