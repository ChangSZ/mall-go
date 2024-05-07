package sms_coupon_history

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
