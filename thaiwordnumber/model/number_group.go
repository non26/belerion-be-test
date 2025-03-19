package model

import "belerionbe/thaiwordnumber/constant"

type NumberGroup struct {
	Value     string
	UnitGroup string
	thaiWord  string
	length    int
}

func NewNumberGroup(value string, unitGroup string) *NumberGroup {
	numberGroup := NumberGroup{
		Value:     value,
		UnitGroup: unitGroup,
		thaiWord:  "",
	}
	numberGroup.length = len(value)
	return &numberGroup
}

// Build the Thai word for the number group with given thaiword merge with thaiword that model have
func (n *NumberGroup) BuildThaiWord(thaiWord string) {
	n.thaiWord = thaiWord + n.UnitGroup
}

func (n *NumberGroup) SetNewThaiWord(number string, unit string) {
	n.thaiWord = number + unit + n.thaiWord
}

// Set the final Thai word for the number group, orverride the thaiword that model have
func (n *NumberGroup) SetFinalThaiWord(thaiWord string) {
	n.thaiWord = thaiWord
}

func (n *NumberGroup) GetLength() int {
	return n.length
}

func (n *NumberGroup) GetThaiWord() string {
	return n.thaiWord
}

// Split the number into groups of at most 6 digits ex. number 7,123,456,789,345,012
// will be split into 3 groups [7123], [456789], [345012]
func (n *NumberGroup) SplitNumber() []*NumberGroup {
	numberOfDigit := 6
	lenNumber := len(n.Value)
	remainder := lenNumber % numberOfDigit
	quotient := lenNumber / numberOfDigit
	unit := n.genrateUnit(quotient, remainder)
	groups := []*NumberGroup{}
	if lenNumber <= numberOfDigit {
		groups = append(groups, &NumberGroup{Value: n.Value, UnitGroup: ""})
		return groups
	}
	if quotient >= 1 {
		if remainder != 0 {
			groups = append(groups, &NumberGroup{Value: n.Value[0:remainder], UnitGroup: unit[0]})
			unit = unit[1:]
		}

		start := remainder
		for i := 0; i < quotient; i++ {
			end := start + numberOfDigit
			groups = append(groups, &NumberGroup{Value: n.Value[start:end], UnitGroup: unit[0]})
			start = end
			unit = unit[1:]
		}
	}
	return groups
}

// Generate the unit for the number group ex. number 7,123,456,789,345,012
// will be split into 3 groups [7123], [456789], [345012]
// the unit will be [ล้าน, ล้าน, ""]
func (n *NumberGroup) genrateUnit(quotient int, remainder int) []string {
	unit := []string{}
	for i := 0; i < quotient; i++ {
		unit = append(unit, constant.MillionReadAs)
	}

	if remainder != 0 {
		unit = append(unit, "")
	} else {
		lastIndex := len(unit) - 1
		unit[lastIndex] = ""
	}
	return unit
}
