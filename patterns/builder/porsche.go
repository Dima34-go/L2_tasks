package builder


type PorsheCar struct {
	Engine int
	Wheel  string
	Color  string
}

func NewPorsheCar() *PorsheCar {
	return &PorsheCar{}
}

func (p *PorsheCar) SetColor() {
	p.Color = "Blue"
}

func (p *PorsheCar) SetEngine() {
	p.Engine = 260
}

func (p *PorsheCar) SetWheel() {
	p.Wheel = "M-220"
}

func (p *PorsheCar) GetCar() Car {
	return Car{
		Engine: p.Engine,
		Wheel:  p.Wheel,
		Color:  p.Color,
	}
}

