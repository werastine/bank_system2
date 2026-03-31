package bank

import "fmt"

type BankAccount struct {
	UserProfile *Profile
	UserCard    string
	IsCreated   bool
	CardMap     map[*Profile]string
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
}

func newProfile(name string, age int, PassportNumber string) *Profile {
	return &Profile{
		name:           name,
		age:            age,
		PassportNumber: PassportNumber,
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
