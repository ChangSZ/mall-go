package sms_flash_promotion_product_relation

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
