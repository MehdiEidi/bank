package currency

import "github.com/MehdiEidi/gods/set"

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

var supportedCurrencies *set.Set[string]

func init() {
	supportedCurrencies = set.New[string]()

	supportedCurrencies.Add(USD)
	supportedCurrencies.Add(EUR)
	supportedCurrencies.Add(CAD)
}

func IsSupported(currency string) bool {
	return supportedCurrencies.Has(currency)
}
