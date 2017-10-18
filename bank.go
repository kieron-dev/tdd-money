package money

type Bank struct {
}

func (b Bank) Reduce(expr Expression, currency string) Money {
	sum := expr.(Sum)
	return sum.Reduce(b, currency)
}
