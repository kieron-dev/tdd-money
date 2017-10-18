package money_test

import (
	"github.com/kieron-pivotal/money"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("can be multiplied by a number", func() {
		fiveBucks := money.NewDollar(5)
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
		fiveBucks := money.NewDollar(5)
		Expect(fiveBucks.Plus(fiveBucks)).To(Equal(money.NewDollar(10)))
	})
})
