package oms_order

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
