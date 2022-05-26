package builder

const(
	LamboCarType = "Lambo"
	PorsheCarType = "Porshe"
)
type Builder interface {
	SetColor()
	SetEngine()
	SetWheel()
	GetCar() Car
}

func getBuilder(builderType string) Builder {
	if builderType == LamboCarType {
		return &LamboCar{}
	}
	if builderType == PorsheCarType {
		return &PorsheCar{}
	}
	return nil
}