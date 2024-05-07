package sms_coupon

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
