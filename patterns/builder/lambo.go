package builder

type LamboCar struct {
	Engine int
	Wheel  string
	Color  string
}

func NewLamboCar() *LamboCar {
	return &LamboCar{}
}

func (l *LamboCar) SetColor() {
	l.Color = "Red"
}

func (l *LamboCar) SetEngine() {
	l.Engine = 200
}

func (l *LamboCar) SetWheel() {
	l.Wheel = "L-270"
}

func (l *LamboCar) GetCar() Car {
	return Car{
		Engine: l.Engine,
		Wheel:  l.Wheel,
		Color:  l.Color,
	}
}


