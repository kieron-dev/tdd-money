package money_test

import (
	"github.com/kieron-pivotal/money"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("can multiply a dollar amount by a number", func() {
		fiveBucks := money.Dollar{Amount: 5}
		product := fiveBucks.Times(2)
		Expect(product).To(Equal(money.Dollar{Amount: 10}))
	})

	It("equals another Dollar with the same amount", func() {
		Expect(money.Dollar{Amount: 5}).To(Equal(money.Dollar{Amount: 5}))
		Expect(money.Dollar{Amount: 5}).ToNot(Equal(money.Dollar{Amount: 6}))
	})
})
