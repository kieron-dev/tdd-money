package money

type Bank struct {
}

func (b Bank) Reduce(expr Expression, currency string) Money {
	return expr.Reduce(b, currency)
}
