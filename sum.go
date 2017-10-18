package money

type Sum struct {
	augend Money
	addend Money
}

func (s Sum) Reduce(bank Bank, currency string) Money {
	augend := bank.Reduce(s.augend, currency)
	addend := bank.Reduce(s.addend, currency)
	return Money{
		currency: augend.currency,
		amount:   augend.amount + addend.amount,
	}
}
