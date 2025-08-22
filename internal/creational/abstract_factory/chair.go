package abstractfactory

type Chair interface {
	SetColor(color string)
	GetColor() string
}

type chair struct {
	color string
}

func (c *chair) SetColor(color string) {
	c.color = color
}

func (c *chair) GetColor() string {
	return c.color
}
