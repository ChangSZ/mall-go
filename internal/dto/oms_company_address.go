package dto

type OmsCompanyAddress struct {
	Id            int64  `json:"id"`
	AddressName   string `json:"addressName"`
	SendStatus    int32  `json:"sendStatus"`
	ReceiveStatus int32  `json:"receiveStatus"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Region        string `json:"region"`
	DetailAddress string `json:"detailAddress"`
}
