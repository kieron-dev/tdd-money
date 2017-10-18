package money_test

import (
	"github.com/kieron-pivotal/money"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("can multiply a dollar amount by a number", func() {
		fiveBucks := money.NewDollar(5)
		product := fiveBucks.Times(2)
		Expect(product).To(Equal(money.NewDollar(10)))
	})

	It("equals another Dollar with the same amount", func() {
		Expect(money.NewDollar(5)).To(Equal(money.NewDollar(5)))
		Expect(money.NewDollar(5)).ToNot(Equal(money.NewDollar(6)))
	})
})
