package oms_order_setting

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
