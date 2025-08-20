package simplefactory

type ak47 struct {
	power int
	name  string
}

func (a *ak47) GetName() string {
	return a.name
}

func (a *ak47) GetPower() int {
	return a.power
}

func (a *ak47) SetName(name string) {
	a.name = name
}

func (a *ak47) SetPower(power int) {
	a.power = power
}

func newAK47() Gun {
	return &ak47{}
}
