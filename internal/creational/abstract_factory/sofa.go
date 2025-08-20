package abstractfactory

type Sofa interface {
	SetColor(color string)
	GetColor() string
}

type sofa struct {
	color string
}

func (s *sofa) SetColor(color string) {
	s.color = color
}

func (s *sofa) GetColor() string {
	return s.color
}
