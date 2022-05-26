package builder


type director struct {
	builder Builder
}

func newDirector(b Builder) *director {
	return &director{
		builder: b,
	}
}
func (dir *director) SetBuilder(b Builder)  {
	dir.builder=b
}
func (dir *director) BuildCar() Car {
	dir.builder.SetEngine()
	dir.builder.SetColor()
	dir.builder.SetWheel()
	return  dir.builder.GetCar()
}


