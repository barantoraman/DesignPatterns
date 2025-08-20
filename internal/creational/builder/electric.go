package builder

type electricCar struct {
	car
}

func newElectricBuilder() Builder {
	return &electricCar{
		car: car{},
	}
}

func (c *electricCar) SetEngine(engine string) {
	c.Engine = engine
}

func (c *electricCar) SetWheel(wheel string) {
	c.Wheel = wheel
}

func (c *electricCar) SetColor(color string) {
	c.Color = color
}

func (c *electricCar) GetCar() *car {
	return &car{
		Engine: c.Engine,
		Wheel:  c.Wheel,
		Color:  c.Color,
	}
}
