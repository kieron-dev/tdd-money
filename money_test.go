package money_test

import (
	"github.com/kieron-pivotal/money"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	Context("Dollar", func() {
		It("can be multiplied by a number", func() {
			fiveBucks := money.NewDollar(5)
			product := fiveBucks.Times(2)
			Expect(product).To(Equal(money.NewDollar(10)))
		})

		It("equals another Dollar with the same amount", func() {
			Expect(money.NewDollar(5)).To(Equal(money.NewDollar(5)))
			Expect(money.NewDollar(5)).ToNot(Equal(money.NewDollar(6)))
		})
	})

	Context("Swiss Franc", func() {
		It("can be multiplied by a number", func() {
			fiveBucks := money.NewFranc(5)
			product := fiveBucks.Times(2)
			Expect(product).To(Equal(money.NewFranc(10)))
		})

		It("equals another Franc with the same amount", func() {
			Expect(money.NewFranc(5)).To(Equal(money.NewFranc(5)))
			Expect(money.NewFranc(5)).ToNot(Equal(money.NewFranc(6)))
		})
	})
})
