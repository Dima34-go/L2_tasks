package factory_method

import "fmt"


const (
	ServerType = "server"
	PersonalComputerType = "personal"
	NotebookType = "notebook"
)
type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s - несущевствующий тип объекта.\n",typeName)
		return nil
	case NotebookType:
		return NewNotebook()
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	}
}