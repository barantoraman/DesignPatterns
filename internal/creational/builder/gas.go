package builder

type gasCar struct {
	car
}

func newGasBuilder() Builder {
	return &gasCar{
		car: car{},
	}
}

func (c *gasCar) SetEngine(engine string) {
	c.Engine = engine
}

func (c *gasCar) SetWheel(wheel string) {
	c.Wheel = wheel
}

func (c *gasCar) SetColor(color string) {
	c.Color = color
}

func (c *gasCar) GetCar() *car {
	return &car{
		Engine: c.Engine,
		Wheel:  c.Wheel,
		Color:  c.Color,
	}
}
