package visitor

import "fmt"

type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}


// Реализация дополнительного функционала для структур

type CalculIncome struct {
      bonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println("Full income including premium",d.Income + d.Income*c.bonusRate/100)
}
func (c CalculIncome) VisitDirector(d Director) {
	fmt.Println("Full income including premium",d.Income + d.Income*c.bonusRate/100)
}

type Vacation struct {
	date string
}

func (v Vacation) VisitDeveloper(d Developer) {
	fmt.Println("Start vacation period",v.date,"Duration:",d.Age*1," days")
}
func (v Vacation) VisitDirector(d Director) {
	fmt.Println("Start vacation period",v.date,"Duration:",d.Age*2," days")
}