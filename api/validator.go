package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mehdieidi/bank/pkg/currency"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if cur, ok := fieldLevel.Field().Interface().(string); ok {
		return currency.IsSupported(cur)
	}
	return false
}
