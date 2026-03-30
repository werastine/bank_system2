package bank

import "fmt"

type BankAccount struct {
	UserProfile *Profile
	UserCard    string
	IsCreated   bool
	CardMap     map[*Profile]string
}

func NewBankAccount() *BankAccount {
	return &BankAccount{
		UserProfile: newProfile(),
		UserCard:    "",
	}
}

type Profile struct {
	balance int
	// Имя и тд будут братся с card
}

func newProfile() *Profile {
	return &Profile{
		balance: 100,
	}
}

func (b *BankAccount) CreateCard(name, passport string) bool {
	card, ok := b.CardCreator(name, passport)
	if !ok {
		fmt.Println("Card is already created")
		return false
	}
	b.UserCard = card
	return true

}
