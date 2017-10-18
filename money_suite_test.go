package money_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMoney(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Money Suite")
}
