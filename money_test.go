package money_test

import (
	"github.com/kieron-pivotal/money"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Money", func() {
	It("can multiply a dollar amount by a number", func() {
		dollars := money.Dollar{Amount: 5}
		dollars.Times(2)
		Expect(dollars.Amount).To(Equal(10))
	})
})
