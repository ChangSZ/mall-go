package oms_order_return_reason

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
