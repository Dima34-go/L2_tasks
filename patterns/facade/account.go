package facade

import "fmt"

type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{
		name: name,
	}
}
func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

