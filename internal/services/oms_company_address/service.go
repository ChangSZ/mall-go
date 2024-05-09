package oms_company_address

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
