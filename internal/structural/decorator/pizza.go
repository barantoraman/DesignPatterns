package decorator

type Pizza interface {
	GetPrice() int
}

type VeggieMania struct{}

func (v *VeggieMania) GetPrice() int {
	return 15
}

type TomatoTopping struct {
	P Pizza
}

func (t *TomatoTopping) GetPrice() int {
	price := t.P.GetPrice()
	return price + 7
}

type CheeseTopping struct {
	P Pizza
}

func (c *CheeseTopping) GetPrice() int {
	price := c.P.GetPrice()
	return price + 10
}
