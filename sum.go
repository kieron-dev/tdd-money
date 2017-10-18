package money

type Sum struct {
	augend Money
	addend Money
}

func (s Sum) Reduce(bank Bank, currency string) Money {
	return Money{
		currency: s.augend.currency,
		amount:   s.augend.amount + s.addend.amount,
	}
}
