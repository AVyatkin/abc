package abc

func New() *Super {
	return &Super{}
}

type Super struct {
	X int64
}

func (s *Super) Hello() string {
	return "hello version 2"
}
