package state

import "fmt"

type noItemState struct {
	vendingMachine *vendingMachine
}

func (h *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock ")
}

func (h *noItemState) addItem(count int) error {
	h.vendingMachine.incrementItemCount(count)
	h.vendingMachine.setState(h.vendingMachine.hasItem)
	return nil
}

func (h *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock ")
}
func (h *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock ")
}