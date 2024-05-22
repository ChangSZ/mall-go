package home

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}
