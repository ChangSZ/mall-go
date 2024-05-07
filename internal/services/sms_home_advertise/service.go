package sms_home_advertise

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
