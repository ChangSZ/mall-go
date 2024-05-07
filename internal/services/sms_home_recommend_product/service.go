package sms_home_recommend_product

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
