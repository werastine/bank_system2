package bank

import (
	"math/rand/v2"
	"strconv"
)

func (b *BankAccount) CardCreator(name, passport string) (string, bool) {
	if !b.IsCreated {
		result := BackEndCardCreator()
		b.IsCreated = true
		return result, true
	} else {
		return "", false
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
	var mod10 string
	var num int
	var cur int
	for i := len(cardNum) - 1; i >= 0; i-- {
		cur = int(cardNum[i] - '0')
		if i%2 != 0 {
			cur *= 2
			if cur > 9 {
				num += cur - 9
			} else {
				num += cur
			}
			continue
		}
		num += cur
	}
	mod10 = strconv.Itoa((10 - (num % 10)) % 10)
	return mod10
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
	cardNumber += "4"
	cardNumber += "044"

	for i := 0; len(cardNumber) < 15; i++ {
		num := RandNumCreation()
		cardNumber += num
	}
	cardNumber += Mod10(cardNumber)

	return cardNumber
}
