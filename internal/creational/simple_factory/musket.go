package simplefactory

type musket struct {
	name  string
	power int
}

func (m *musket) GetName() string {
	return m.name
}

func (m *musket) GetPower() int {
	return m.power
}

func (m *musket) SetName(name string) {
	m.name = name
}

func (m *musket) SetPower(power int) {
	m.power = power
}

func newMusket() Gun {
	return &musket{}
}
