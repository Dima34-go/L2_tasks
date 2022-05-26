package visitor

import "fmt"

type Developer struct {
	FirstName string
	LastName string
	Income int
	Age int
}


func (d Developer) FullName() {
	fmt.Printf("Full name: %s %s Age: %d years, Income: %d$, Developer \n",d.FirstName,d.LastName,d.Age,d.Income)
}

func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}
