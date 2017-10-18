package money

type Bank struct {
	rates map[[2]string]int
}

func NewBank() Bank {
	bank := Bank{}
	bank.rates = make(map[[2]string]int)
	return bank
}

func (b Bank) Reduce(expr Expression, currency string) Money {
	return expr.Reduce(b, currency)
}

func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[[2]string{from, to}] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[[2]string{from, to}]
}
