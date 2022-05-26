package chainOfResponsibility

func MainChainResponsibility()  {
	//Задание структуры
	device := Device{Name: "Device-1"}
	updateSvc := UpdateDataService{Name: "Update-1"}
	dataSvc := DataService{}
	device.SetNext(&updateSvc)
	updateSvc.SetNext(&dataSvc)

	//Пример работы
	data := Data{
		GetSource:    false,
		UpdateSource: false,
	}
	device.Execute(&data)
}