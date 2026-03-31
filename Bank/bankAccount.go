package bank

import (
	"fmt"
)

type BankAccount struct {
	UserProfile *Profile
	UserCard    string
	IsCreated   bool
	// CardMap     map[]string
}

func NewBankAccount(name string, age int, PassportNumber string) *BankAccount {
	return &BankAccount{
		UserProfile: newProfile(name, age, PassportNumber),
		UserCard:    "",
	}
}

type Profile struct {
	name           string
	age            int
	PassportNumber string
	UserID         string
}

func newProfile(name string, age int, PassportNumber string) *Profile {
	return &Profile{
		name:           name,
		age:            age,
		PassportNumber: PassportNumber,
		UserID:         IDGenerator(),
	}
}

func (b *BankAccount) CreateCard() bool { // There i have to change code
	card, ok := b.CardCreator(b.UserProfile.name, b.UserProfile.PassportNumber, b.UserProfile.age)
	if !ok {
		fmt.Println("Card is already created")
		return false
	}
	b.UserCard = card
	return true

}

func IDGenerator() string {
	var ID string
	for i := 0; i < 12; i++ {
		num := RandNumCreation()
		ID += num
	}
	return ID
}
