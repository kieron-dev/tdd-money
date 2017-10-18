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
		bank := money.Bank{}
		reduced := bank.Reduce(sum, "USD")
		Expect(reduced).To(Equal(money.NewDollar(10)))
	})

	It("returns a Sum when added", func() {
		sum := fiveBucks.Plus(fiveBucks)
		Expect(sum).To(BeAssignableToTypeOf(money.Sum{}))
	})

	It("reduces to a Money", func() {
		bank := money.Bank{}
		reduced := bank.Reduce(fiveBucks, "USD")
		Expect(reduced).To(Equal(fiveBucks))
	})
})
