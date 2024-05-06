package dto

import "time"

type CmsSubjectProductRelation struct {
	Id        int64 `json:"id"`        //
	SubjectId int64 `json:"subjectId"` //
	ProductId int64 `json:"productId"` //
}

type CmsSubject struct {
	Id              int64     `json:"id"`
	CategoryId      int64     `json:"categoryId"`
	Title           string    `json:"title"`
	Pic             string    `json:"pic"`
	ProductCount    int32     `json:"productCount"`
	RecommendStatus int32     `json:"recommendStatus"`
	CreateTime      time.Time `json:"createTime"`
	CollectCount    int32     `json:"collectCount"`
	ReadCount       int32     `json:"readCount"`
	CommentCount    int32     `json:"commentCount"`
	AlbumPics       string    `json:"albumPics"`
	Description     string    `json:"description"`
	ShowStatus      int32     `json:"showStatus"`
	Content         string    `json:"content"`
	ForwardCount    int32     `json:"forwardCount"`
	CategoryName    string    `json:"categoryName"`
}
