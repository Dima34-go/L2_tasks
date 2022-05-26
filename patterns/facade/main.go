package facade

import (
	"fmt"
	"log"
)


func MainFacade() {
	fmt.Println()
	walletFacade := NewWalletFacade("abc", "1234")
	fmt.Println()
	err := walletFacade.addMoneyToWallet("abc", "1234", 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", "1234", 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}