package GOL2014_test

import (
	. "github.com/adamkalnas/GOL2014"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GOL", func() {
	It("should be a short story", func() {
		Expect(DasBoot()).To(Equal(1))
	})
})
