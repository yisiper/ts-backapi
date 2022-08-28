package processOrder_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProcessOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProcessOrder Suite")
}
