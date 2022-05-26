package facade

import "fmt"

type securityCode struct {
	scCode string
}

func newSecurityCode(scCode string) *securityCode {
	return &securityCode{scCode: scCode}
}

func (s *securityCode) checkCode(incomingCode string) error {
	if s.scCode != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
