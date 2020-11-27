//go:generate go run github.com/mojadev/camunda-go-delegate/generate --config ./delegate.yaml --output ../main.go

package delegates

import (
	"errors"
	"github.com/mojadev/camunda-go-delegate/api"
)

type CurrencyContainer struct {
	Amount int
}

func ConvertCurrency(ctx *api.ExecutionContext) (api.ExecutionResult, error)  {
	result := api.CreateExecutionResult()
	currency, ok := ctx.Variables["currency"].(CurrencyContainer)

	if ok == false {
		return result, errors.New("could not read currency")
	}

	currency.Amount = currency.Amount * 2
	result.Variables["currency"] = currency
	return result, nil
}
