package domain_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

var _ = Describe("Domain", func() {
	p, _ := domain.NewPromise(domain.GenerateID(), 1, "desc", nil)

	var newPromise domain.Promise

	BeforeEach(func() {
		newPromise = *p
	})

	Describe("New promise", func() {
		Context("With status", func() {
			It("should be a `draft`", func() {
				Expect(newPromise.Status).To(Equal("draft"))
			})
		})

		Context("With timestamp", func() {
			It("CreatedAt should be equal UpdatedAt", func() {
				Expect(newPromise.CreatedAt).To(Equal(newPromise.UpdatedAt))
			})
		})
	})
})
