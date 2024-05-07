package sms_home_brand

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
