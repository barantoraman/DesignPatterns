package builder

type Director interface {
	SetWheel(wheel string)
	SetEngine(engine string)
	SetColor(color string)
	Build() *car
}

type director struct {
	builder Builder
}

func (d *director) Build() *car {
	return d.builder.GetCar()
}

func (d *director) SetColor(color string) {
	d.builder.SetColor(color)
}

func (d *director) SetEngine(engine string) {
	d.builder.SetEngine(engine)
}

func (d *director) SetWheel(wheel string) {
	d.builder.SetWheel(wheel)
}

func NewDirector(builder Builder) Director {
	if builder == nil {
		return nil
	}
	return &director{
		builder: builder,
	}
}
