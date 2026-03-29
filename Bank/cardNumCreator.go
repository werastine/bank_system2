package bank

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func (b *BankAccount) CardCreator(name, passport string) (string, error) {
	if !b.IsCreated {
		result := BackEndCardCreator()
		//Card creation process where I gonna use mod-10 alg.
		b.IsCreated = true
		return result, nil
	} else {
		return "", fmt.Errorf("Action blocked, You cannot create card again")
	}
	// implement mod-10 algorithm the same as I've made in previous one

}

func RandNumCreation() string {
	randNum := rand.IntN(9)
	res := strconv.Itoa(randNum)
	return res

}

func Mod10(cardNum string) string {
	// there will be realisation of mod10 alg

	return "last digit"
}

func BackEndCardCreator() string {
	// so there i will be full implementation of card num creation
	// Card number Contain 16 digits:
	// 1: first digit is visa(4) or mastercard(5), my bank work with visa so: 4
	// 2-4: The next 3 digits is our bank locan number (044)
	// 5-15: The rest(except last digit) is fully random number (0-9)
	// 16: last digit is mode-10 one, which we can know only after creation of previous 15 digits
	// --------------------------------------------------------------------------------------------- //
	var cardNumber string
	cardNumber += "044"

	for i := 0; len(cardNumber) < 16; i++ {
		num := RandNumCreation()
		cardNumber += num
	}

	cardNumber += Mod10(cardNumber)

	return cardNumber
}
