package sms_flash_promotion_session

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
