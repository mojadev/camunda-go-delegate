package steps_test

import (
	"github.com/mojadev/camunda-go-delegate/pkg/api"
	"github.com/mojadev/camunda-go-delegate-example/steps"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example", func() {
	Context("convertCurrency", func() {
		When("a currency variable is set", func() {
			var variables api.VariableScope = make(api.VariableScope);
			variables["currency"] = steps.CurrencyContainer{Amount: 100}
			context := api.ExecutionContext{
				Variables: variables,
			}
			result, err := steps.ConvertCurrency(&context)
			It("returns a variableScope with the converted currency", func() {
				Expect(err).To(BeNil())
				Expect(result.Variables["currency"]).To(BeEquivalentTo(steps.CurrencyContainer{Amount: 200}))
			})
		})
	})
})
