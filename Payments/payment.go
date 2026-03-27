package payments

// There i will implement interface for different payment methods

type Payer interface {
	pay(sum int)
	cancel(id int)
}

type Payment struct {
	transactionID          int
	tranasactionSum        int
	transactionStatus      bool
	transactionDescription string
}

func newPayment() *Payment {
	return &Payment{}
}

func (p *Payment) pay(sum int, desctiption string) {
	p.tranasactionSum = sum
	p.transactionDescription = desctiption
	// Than i need to check current balance
	// to - sum from the balance
	//
}

// I haven't finished it. Firstly I'd want to
