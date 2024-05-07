package sms_home_recommend_subject

// SmsHomeRecommendSubject 首页推荐专题表
//
//go:generate gormgen -structs SmsHomeRecommendSubject -input .
type SmsHomeRecommendSubject struct {
	Id              int64  //
	SubjectId       int64  //
	SubjectName     string //
	RecommendStatus int32  //
	Sort            int32  //
}
