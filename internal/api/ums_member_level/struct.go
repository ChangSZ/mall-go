package ums_member_level

type UmsMemberLevel struct {
	Id                    int64   `json:"id"`
	Name                  string  `json:"name"`
	GrowthPoint           int32   `json:"growthPoint"`
	DefaultStatus         int32   `json:"defaultStatus"`
	FreeFreightPoint      float64 `json:"freeFreightPoint"`
	CommentGrowthPoint    int32   `json:"commentGrowthPoint"`
	PriviledgeFreeFreight int32   `json:"priviledgeFreeFreight"`
	PriviledgeSignIn      int32   `json:"priviledgeSignIn"`
	PriviledgeComment     int32   `json:"priviledgeComment"`
	PriviledgePromotion   int32   `json:"priviledgePromotion"`
	PriviledgeMemberPrice int32   `json:"priviledgeMemberPrice"`
	PriviledgeBirthday    int32   `json:"priviledgeBirthday"`
	Note                  string  `json:"note"`
}
