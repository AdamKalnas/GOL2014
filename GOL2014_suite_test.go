package GOL2014_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGOL2014(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GOL2014 Suite")
}
