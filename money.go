package money

type Money struct {
	currency string
	amount   int
}

type Expression interface {
	Reduce(bank Bank, currency string) Money
	Plus(expr Expression) Expression
}

func NewDollar(amount int) Money {
	return Money{currency: "USD", amount: amount}
}

func NewFranc(amount int) Money {
	return Money{currency: "CHF", amount: amount}
}

func (d Money) Times(multiplier int) Money {
	return Money{currency: d.currency, amount: d.amount * multiplier}
}

func (d Money) Plus(addend Expression) Expression {
	return Sum{augend: d, addend: addend}
}

func (d Money) Reduce(bank Bank, currency string) Money {
	rate := bank.Rate(d.currency, currency)
	return Money{currency: currency, amount: d.amount / rate}
}
