package money

type Dollar struct {
	amount int
}

type Franc struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}

func NewFranc(amount int) Franc {
	return Franc{amount: amount}
}

func (f Franc) Times(multiplier int) Franc {
	return Franc{amount: f.amount * multiplier}
}
