package delegates_test

import (
	"github.com/mojadev/camunda-go-delegate/api"
	"github.com/mojadev/camunda-go-delegate/example/delegates"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example", func() {
	Context("convertCurrency", func() {
		When("a currency variable is set", func() {
			var variables api.VariableScope = make(api.VariableScope);
			variables["currency"] = delegates.CurrencyContainer{Amount: 100}
			context := api.ExecutionContext{
				Variables: variables,
			}
			result, err := delegates.ConvertCurrency(&context)
			It("returns a variableScope with the converted currency", func() {
				Expect(err).To(BeNil())
				Expect(result.Variables["currency"]).To(BeEquivalentTo(delegates.CurrencyContainer{Amount: 200}))
			})
		})
	})
})
