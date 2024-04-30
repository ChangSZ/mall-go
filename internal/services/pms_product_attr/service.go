package pms_product_attr

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
