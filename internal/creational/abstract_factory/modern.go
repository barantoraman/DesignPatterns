package abstractfactory

type modern struct{}
type modernChair struct{ chair }
type modernSofa struct{ sofa }

func (m *modern) MakeChair() Chair {
	return &modernChair{
		chair: chair{},
	}
}

func (m *modern) MakeSofa() Sofa {
	return &modernSofa{
		sofa: sofa{},
	}
}
