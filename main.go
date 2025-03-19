package main

import (
	"belerionbe/thaiwordnumber"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	thaiWordNumber := thaiwordnumber.NewThaiWordNumber()
	for _, input := range inputs {
		fmt.Println(input)
		// convert decimal to thai text (baht) and print the result here
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(input)
		println(thaiWord)
	}
}
