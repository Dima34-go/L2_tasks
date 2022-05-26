package visitor

func MainVisitor()  {
	backend := Developer{"Bob", "Bilbo", 1000, 32}
	boss := Director{"Bob", "Baggins", 2000, 40}
	backend.FullName()
	backend.Accept(CalculIncome{20})
	backend.Accept(Vacation{"17.05.2022"})
	boss.FullName()
	boss.Accept(CalculIncome{10})
	boss.Accept(Vacation{"12.01.2023"})
}
