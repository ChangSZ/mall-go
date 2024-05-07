package sms_home_new_product

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
