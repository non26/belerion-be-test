package model

import (
	"belerionbe/thaiwordnumber/constant"
	"strings"

	"github.com/shopspring/decimal"
)

type Number struct {
	WholeNumber   *NumberGroup
	DecimalNumber *NumberGroup
}

// this function will split number into 2 part, whole and decimal part
// example: 1234.50 will be split into whole part = 1234 and decimal part = 50
func NewNumberWith(number decimal.Decimal) *Number {
	numberString := number.String()
	numberArray := strings.Split(numberString, ".") // split whole part at index 0 and decimal part at index 1
	lenNumberArray := len(numberArray)
	if lenNumberArray == 2 {
		lenDecimal := len(numberArray[1])
		if lenDecimal == 1 {
			numberArray[1] = numberArray[1] + "0" // case where decimal part is 1 digit add 0 to the end as 10
		}
		if lenDecimal == 2 && numberArray[1][0] == '0' { // case where decimal part is 2 digit and first digit is 0 as 05
			numberArray[1] = numberArray[1][1:] // remove the first digit so that it is 5 only
		}
	}
	var newNumber *Number
	if lenNumberArray == 1 { // number 1234
		newNumber = newNumberWithOnlyWholePart(numberArray[0])
	} else if lenNumberArray == 2 {
		if numberArray[0] == "0" { // 0.50
			newNumber = newNumberWithOnlyDecimalPart(numberArray[1])
		} else if numberArray[1] == "0" || numberArray[1] == "00" { // 1234.0 or 1234.00
			newNumber = newNumberWithOnlyWholePart(numberArray[0])
		} else {
			newNumber = newNumberWith(numberArray[0], numberArray[1]) // 1234.50
		}
	}
	if newNumber.DecimalNumber == nil {
		newNumber.WholeNumber.UnitGroup = constant.Baht + constant.Total // บาทถ้วน
	}

	return newNumber
}

func newNumberWithOnlyWholePart(value string) *Number {
	number := &Number{
		WholeNumber:   &NumberGroup{Value: value, UnitGroup: constant.Baht},
		DecimalNumber: nil,
	}
	return number
}

func newNumberWithOnlyDecimalPart(value string) *Number {
	number := &Number{
		WholeNumber:   nil,
		DecimalNumber: &NumberGroup{Value: value, UnitGroup: constant.Stang},
	}
	return number
}

func newNumberWith(whole_value string, decimal_value string) *Number {
	number := &Number{
		WholeNumber:   &NumberGroup{Value: whole_value, UnitGroup: constant.Baht},
		DecimalNumber: &NumberGroup{Value: decimal_value, UnitGroup: constant.Stang},
	}
	return number
}

func (n *Number) GetThaiWord() string {
	thaiWord := ""
	if n.WholeNumber != nil {
		thaiWord = n.WholeNumber.GetThaiWord() + n.WholeNumber.UnitGroup
	}
	if n.DecimalNumber != nil {
		thaiWord = thaiWord + n.DecimalNumber.GetThaiWord() + n.DecimalNumber.UnitGroup
	}
	return thaiWord
}
