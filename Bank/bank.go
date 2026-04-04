package bank

type Bank struct {
	CardStorage map[string]struct{}
	IDStorage   map[string]struct{}
}

func NewBank() *Bank {
	return &Bank{
		CardStorage: make(map[string]struct{}),
		IDStorage:   make(map[string]struct{}),
	}
}
