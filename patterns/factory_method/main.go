package factory_method

func MainFactory()  {
	var types = []string{PersonalComputerType,ServerType,ServerType,NotebookType,"wrong_type"}
	for _,typeName := range types {
		cmp := New(typeName)
		if cmp != nil {
			cmp.PrintDetails()
		}
	}
}