package currency

import "github.com/MehdiEidi/gods/set"

func IsSupported(currency string) bool {
	supportedCurrencies := set.New[string]()

	supportedCurrencies.Add("USD")
	supportedCurrencies.Add("EUR")
	supportedCurrencies.Add("CAD")

	return supportedCurrencies.Has(currency)
}
