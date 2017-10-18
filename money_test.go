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
		Expect(product.Amount).To(Equal(10))
	})
})
