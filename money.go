package money

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(multiplier int) {
	d.Amount = d.Amount * multiplier
}
