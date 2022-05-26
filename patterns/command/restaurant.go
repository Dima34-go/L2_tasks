package command

import "fmt"

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

//создание нового ресторана

func NewRestaurant() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}


func (r *Restaurant) MakePizza(count int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		count:          count,
	}
}

func (r *Restaurant) MakeSalad(count int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		count:         count,
	}
}

func (r *Restaurant) CleanDishes(count int) Command {
	return &CleanDishesCommand{
		count: count,
		restaurant: r,
	}
}
//команда создания пиццы

type MakePizzaCommand struct{
	count     int
	restaurant *Restaurant
}

func (m *MakePizzaCommand) execute() {
	m.restaurant.CleanedDishes -= m.count
	fmt.Println("made",m.count,"pizzas")
}

//команда создания салата

type MakeSaladCommand struct{
	count     int
	restaurant *Restaurant
}

func (m *MakeSaladCommand) execute() {
	m.restaurant.CleanedDishes -= m.count
	fmt.Println("made",m.count,"salads")
}

//очистка посуды

type CleanDishesCommand  struct{
	count  int
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	c.restaurant.CleanedDishes += c.count
	fmt.Println("cleaned",c.count,"dishes")
}
