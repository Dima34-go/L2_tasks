package state

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h *hasItemState) addItem(count int) error {
	h.vendingMachine.incrementItemCount(count)
	fmt.Printf("%d items added\n", count)
	return nil
}

func (h *hasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("No item present ")
	}
	fmt.Printf("Item requested\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first ")
}

func (h *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first ")
}


