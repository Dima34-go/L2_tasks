package builder

import "fmt"

func MainBuilder() {

	//создание интерфейсов строителя
	bL:=getBuilder(LamboCarType)
	bP:=getBuilder(PorsheCarType)

	//Создан dir, с указанием создаваемой машины
	dir:=newDirector(bL)

	//Создание машины
	lCar:=dir.BuildCar()
	fmt.Println(lCar)

	//Смена типа машины
	dir.SetBuilder(bP)
	pCar:=dir.BuildCar()
	fmt.Println(pCar)
}