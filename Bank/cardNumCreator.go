package bank

import "fmt"

func (b *BankAccount) CardCreator(name, passport string) (string, error) {
	if !b.IsCreated {
		//Card creation process where i gonna use mod-10 alg.
		b.IsCreated = true
		return "cardnum", nil
	} else {
		return "", fmt.Errorf("Action blocked, You cannot create card again")
	}
	// implement mod-10 algorithm the same as I've made in previous one

}
