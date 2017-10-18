package money_test

import (
	"github.com/kieron-pivotal/money"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	fiveBucks money.Money
)

var _ = Describe("Money", func() {
	BeforeEach(func() {
		fiveBucks = money.NewDollar(5)
	})

	It("can be multiplied by a number", func() {
		product := fiveBucks.Times(2)
		Expect(product).To(Equal(money.NewDollar(10)))

		fiveFrancs := money.NewFranc(5)
		product = fiveFrancs.Times(2)
		Expect(product).To(Equal(money.NewFranc(10)))
	})

	It("equals another Dollar with the same amount", func() {
		Expect(money.NewDollar(5)).To(Equal(money.NewDollar(5)))
		Expect(money.NewDollar(5)).ToNot(Equal(money.NewDollar(6)))
		Expect(money.NewDollar(5)).ToNot(Equal(money.NewFranc(5)))
	})

	It("can be added to another money", func() {
		sum := fiveBucks.Plus(fiveBucks)
		bank := money.NewBank()
		reduced := bank.Reduce(sum, "USD")
		Expect(reduced).To(Equal(money.NewDollar(10)))
	})

	It("returns a Sum when added", func() {
		sum := fiveBucks.Plus(fiveBucks)
		Expect(sum).To(BeAssignableToTypeOf(money.Sum{}))
	})

	It("reduces to a Money", func() {
		bank := money.NewBank()
		reduced := bank.Reduce(fiveBucks, "USD")
		Expect(reduced).To(Equal(fiveBucks))
	})

	It("reduces to a different currency using exchange rate", func() {
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		tenFrancs := money.NewFranc(10)
		reduced := bank.Reduce(tenFrancs, "USD")
		Expect(reduced).To(Equal(fiveBucks))
	})

	It("can add two currencies and reduce using rate", func() {
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		tenFrancs := money.NewFranc(10)
		sum := fiveBucks.Plus(tenFrancs)
		reduced := bank.Reduce(sum, "USD")
		Expect(reduced).To(Equal(money.NewDollar(10)))
	})
})

var _ = Describe("Bank", func() {
	It("length 2 string arrays with same vals are the same", func() {
		val1 := [2]string{"asdf", "foo"}
		val2 := [2]string{"asdf", "foo"}
		Expect(val1).To(Equal(val2))
	})

	It("can retrieve a rate", func() {
		bank := money.NewBank()
		bank.AddRate("CHF", "USD", 2)
		rate := bank.Rate("CHF", "USD")
		Expect(rate).To(Equal(2))
	})

	It("gives 1 for a same currency rate", func() {
		bank := money.NewBank()
		Expect(bank.Rate("ABC", "ABC")).To(Equal(1))
	})
})
