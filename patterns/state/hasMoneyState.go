package state

import "fmt"

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h *hasMoneyState) addItem(i int) error {
	return fmt.Errorf("Item dispense in progress ")
}

func (h *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress ")
}

func (h *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Money has already been deposited ")
}

func (h *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}
