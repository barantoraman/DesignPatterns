package abstractfactory

type victorian struct{}
type victorianChair struct{ chair }
type victorianSofa struct{ sofa }

func (v *victorian) MakeChair() Chair {
	return &victorianChair{
		chair: chair{},
	}
}

func (v *victorian) MakeSofa() Sofa {
	return &victorianSofa{
		sofa: sofa{},
	}
}

func (v *victorianChair) SetColor(color string) {
	v.chair.color = "victorian color"
}
