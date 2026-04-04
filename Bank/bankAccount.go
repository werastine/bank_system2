package bank

import (
	"fmt"
)

type BankAccount struct {
	UserProfile *Profile
	UserCard    string
	IsCreated   bool
	CardStorage map[string]struct{}
	IDStorage   map[string]struct{}
	UserID      string
}

func NewBankAccount(name string, age int, PassportNumber string) *BankAccount {
	acc := &BankAccount{
		UserCard:    "",
		CardStorage: make(map[string]struct{}),
		IDStorage:   make(map[string]struct{}),
	}
	acc.UserProfile = newProfile(name, age, PassportNumber, acc.IDStorage)
	return acc
}

type Profile struct {
	name           string
	age            int
	PassportNumber string
	UserID         string
}

func newProfile(name string, age int, PassportNumber string, IDstorage map[string]struct{}) *Profile {
	return &Profile{
		name:           name,
		age:            age,
		PassportNumber: PassportNumber,
		UserID:         IDGenerator(IDstorage),
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

func IDGenerator(storage map[string]struct{}) string {

	for {
		var ID string
		for i := 0; i < 12; i++ {
			ID += RandNumCreation()
		}
		if _, ok := storage[ID]; ok {
			{
				continue
			}
		}
		storage[ID] = struct{}{}
		return ID
	}
}
