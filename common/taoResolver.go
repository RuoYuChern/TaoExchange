package common

type TaoResolver struct {
	role int
}

func NewTao(role int) *TaoResolver {
	s := new(TaoResolver)
	s.role = role
	return s
}

func (s *TaoResolver) connect(url string) {
	//GrpcResolver.update()
}
