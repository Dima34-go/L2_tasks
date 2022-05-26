package factory_method

import "fmt"

type Notebook struct {
	Type   string
	Core   int
	Memory int
	Monitor bool
}

func NewNotebook() Computer {
	return &Notebook{
		Type:   NotebookType,
		Core:   16,
		Memory: 256,
		Monitor: true,
	}
}

func (s *Notebook) GetType() string {
	return s.Type
}

func (s *Notebook) PrintDetails() {
	fmt.Printf("%s : Core: [%d], Memory: [%d], Monitor: [%v]\n",s.Type,s.Core,s.Memory,s.Monitor)
}