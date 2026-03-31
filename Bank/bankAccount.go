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
	IDStorage      map[string]struct{}
}

func newProfile(name string, age int, PassportNumber string) *Profile {
	p := &Profile{
		name:           name,
		age:            age,
		PassportNumber: PassportNumber,
		IDStorage:      make(map[string]struct{}),
	}
	p.UserID = p.IDGenerator()
	return p
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

func (p *Profile) IDGenerator() string {

	for {
		var ID string
		for i := 0; i < 12; i++ {
			ID += RandNumCreation()
		}
		if _, ok := p.IDStorage[ID]; ok {
			{
				continue
			}
		}
		p.IDStorage[ID] = struct{}{}
		return ID
	}
}
