package dto

type SmsHomeRecommendSubject struct {
	Id              int64  `json:"id"`
	SubjectId       int64  `json:"subjectId"`
	SubjectName     string `json:"subjectName"`
	RecommendStatus int32  `json:"recommendStatus"`
	Sort            int32  `json:"sort"`
}
