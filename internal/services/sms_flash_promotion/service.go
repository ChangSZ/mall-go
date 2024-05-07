package sms_flash_promotion

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
