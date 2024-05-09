package oms_order_return_apply

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
