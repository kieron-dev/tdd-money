package money

type Sum struct {
	augend Expression
	addend Expression
}

func (s Sum) Plus(expr Expression) Expression {
	return Sum{augend: s, addend: expr}
}

func (s Sum) Times(multiplier int) Expression {
	return Sum{
		augend: s.augend.Times(multiplier),
		addend: s.addend.Times(multiplier),
	}
}

func (s Sum) Reduce(bank Bank, currency string) Money {
	augend := bank.Reduce(s.augend, currency)
	addend := bank.Reduce(s.addend, currency)
	return Money{
		currency: augend.currency,
		amount:   augend.amount + addend.amount,
	}
}
