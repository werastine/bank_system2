package bank

type BankAccount struct {
	UserProfile *Profile
	UserCard    string
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

func (b BankAccount) CreateCard(name, passport string) {
	card := CardCreator(name, passport)
	b.UserCard = card
}
