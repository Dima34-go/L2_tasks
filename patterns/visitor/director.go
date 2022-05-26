package visitor

import "fmt"

type Director struct {
	FirstName string
	LastName string
	Income int
	Age int
}


func (d Director) FullName() {
	fmt.Printf("Full name: %s %s Age: %d years, Income: %d$, Director Company \n",d.FirstName,d.LastName,d.Age,d.Income)
}

func (d Director) Accept(v Visitor) {
	v.VisitDirector(d)
}