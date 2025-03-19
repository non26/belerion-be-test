package factorization

import (
	"belerionbe/factorization/model"

	"github.com/shopspring/decimal"
)

func NewFactorization(number string) decimal.Decimal {
	val, _ := decimal.NewFromString(number)
	return val
}

func Factorize(number decimal.Decimal) []model.Factor {
	factors := []model.Factor{}
	var power int64 = 1
	lenInt64 := int64(len(number.String()))

	for power <= lenInt64 { // number = 199
		remainder := number.Mod(decimal.NewFromInt(10).Pow(decimal.NewFromInt(power))) // remainder will be integer ex 199mod10 = 9
		powerAtThatDigit := decimal.NewFromInt(10).Pow(decimal.NewFromInt(power - 1))
		n := remainder.Div(powerAtThatDigit) // n will be integer ex 9/10^(1-1)
		factor := model.NewFactor(n.String(), powerAtThatDigit.String())
		factors = append(factors, factor)
		number = number.Sub(remainder) // 199 - 9 = 190
		power++
	}
	return factors
}
