package sms_home_recommend_subject

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
