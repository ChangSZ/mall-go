package dto

type CmsPrefrenceAreaProductRelation struct {
	Id              int64 `json:"id"`              //
	PrefrenceAreaId int64 `json:"prefrenceAreaId"` //
	ProductId       int64 `json:"productId"`       //
}

type CmsPrefrenceArea struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	SubTitle   string `json:"subtitle"`
	Pic        []byte `json:"pic"`
	Sort       int32  `json:"sort"`
	ShowStatus int32  `json:"showStatus"`
}
