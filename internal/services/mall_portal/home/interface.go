package home

var _ Service = (*service)(nil)

type Service interface {
	i()
}
