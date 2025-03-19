package model

import "belerionbe/thaiwordnumber/constant"

type NumberInThaiWord struct { // change name to NumberInThaiWord
	numberMap map[string]string
}

func NewNumberInThaiWord() *NumberInThaiWord {
	n := &NumberInThaiWord{
		numberMap: map[string]string{},
	}
	n.numberMap[constant.Zero] = constant.ZeroReadAs
	n.numberMap[constant.One] = constant.OneReadAs
	n.numberMap[constant.Two] = constant.TwoReadAs
	n.numberMap[constant.Three] = constant.ThreeReadAs
	n.numberMap[constant.Four] = constant.FourReadAs
	n.numberMap[constant.Five] = constant.FiveReadAs
	n.numberMap[constant.Six] = constant.SixReadAs
	n.numberMap[constant.Seven] = constant.SevenReadAs
	n.numberMap[constant.Eight] = constant.EightReadAs
	n.numberMap[constant.Nine] = constant.NineReadAs
	n.numberMap[constant.Ten] = constant.TenReadAs
	n.numberMap[constant.Twenty] = constant.TwentyReadAs
	return n
}

// Get the Thai word for the given value and power
func (n *NumberInThaiWord) Get(last_set_of_number bool, value string, power string) string {
	result := n.rule1(value, power)
	if result != nil {
		return *result
	}

	result = n.rule2(value, power)
	if result != nil {
		return *result
	}

	result = n.rule3(last_set_of_number, value, power)
	if result != nil {
		return *result
	}

	return n.numberMap[value]
}

// Get the Thai word for the given value as 1 and power as 10
func (n *NumberInThaiWord) rule1(value string, power string) *string {
	if value == constant.One && power == constant.Ten {
		result := ""
		return &result
	}
	return nil
}

// Get the Thai word for the given value as 2 and power as 10
func (n *NumberInThaiWord) rule2(value string, power string) *string {
	if value == constant.Two && power == constant.Ten {
		result := n.numberMap[constant.Twenty]
		return &result
	}
	return nil
}

// Get the Thai word for the given value as 1 and power as 1 and last_set_of_number is true
// this rule will give pronounce the last digit as 1 as เอ็ด if the number is more that 2 digit
func (n *NumberInThaiWord) rule3(last_set_of_number bool, value string, power string) *string {
	if last_set_of_number && value == constant.One && power == constant.One {
		return &constant.LastDigitOneReadAs
	}
	return nil
}
