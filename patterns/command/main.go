package command

import "fmt"

func MainCommand() {

	r := NewRestaurant()

	// create the list of tasks to be executed
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(10),
		r.MakePizza(4),
		r.CleanDishes(10),
	}

	// create the cooks that will execute the tasks
	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	//распределение задач по поварам

	for i, task := range tasks {
		// Using the modulus of the current task index, we can
		// alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.commands = append(cook.commands, task)
	}

	//выполнение задач поварами

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}