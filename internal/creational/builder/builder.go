package builder

type Builder interface {
	SetEngine(engine string)
	SetWheel(wheel string)
	SetColor(color string)
	GetCar() *car
}

func NewBuilder(builderType string) Builder {
	switch builderType {
	case "electric":
		return newElectricBuilder()
	case "gas":
		return newGasBuilder()
	default:
		return nil
	}
}
