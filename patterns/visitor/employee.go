package visitor

type Employee interface {
	FullName()
	Accept(Visitor)
}