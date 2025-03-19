package thaiwordnumber

import (
	"belerionbe/factorization"
	"belerionbe/thaiwordnumber/model"

	"github.com/shopspring/decimal"
)

type thaiWordNumber struct {
	thaiWordNumber *model.NumberInThaiWord
	magnitude      *model.Magnitude
}

func NewThaiWordNumber() *thaiWordNumber {
	return &thaiWordNumber{
		thaiWordNumber: model.NewNumberInThaiWord(),
		magnitude:      model.NewMagnitude(),
	}
}

func (t *thaiWordNumber) GetThaiWordNumberOf(number decimal.Decimal) string {
	newNumber := model.NewNumberWith(number)
	var groups []*model.NumberGroup
	if newNumber.WholeNumber != nil {
		groups = newNumber.WholeNumber.SplitNumber()
		wholeNumberReadsThaiWord := t.factorize(newNumber.WholeNumber.Value, groups)
		newNumber.WholeNumber.SetFinalThaiWord(wholeNumberReadsThaiWord)
	}
	if newNumber.DecimalNumber != nil {
		groups = newNumber.DecimalNumber.SplitNumber()
		decimalNumberReadsThaiWord := t.factorize(newNumber.DecimalNumber.Value, groups)
		newNumber.DecimalNumber.SetFinalThaiWord(decimalNumberReadsThaiWord)
	}
	finalWord := newNumber.GetThaiWord()
	return finalWord
}

// Factorize the number into Thai word
func (t *thaiWordNumber) factorize(value string, setOfNumberGroup []*model.NumberGroup) string {
	finalThaiWord := ""
	lengthGroup := len(setOfNumberGroup)
	for idxg, numberGroup := range setOfNumberGroup {
		if len(numberGroup.Value) == 1 {
			thaiWord := t.thaiWordNumber.Get(false, numberGroup.Value, "")
			numberGroup.BuildThaiWord(thaiWord)
			finalThaiWord = finalThaiWord + numberGroup.GetThaiWord()
			continue
		}

		val := factorization.NewFactorization(numberGroup.Value)
		factors := factorization.Factorize(val)
		for idxf, factor := range factors {
			lastSetOfNumber := (lengthGroup == idxg+1) && (idxf == 0)
			if factor.Number == "0" {
				continue
			}
			numberInThaiWord := t.thaiWordNumber.Get(lastSetOfNumber, factor.Number, factor.Power)
			UnitInThatWord := t.magnitude.Get(factor.Power)
			numberGroup.SetNewThaiWord(numberInThaiWord, UnitInThatWord)

		}
		numberGroup.BuildThaiWord(numberGroup.GetThaiWord())
		finalThaiWord = finalThaiWord + numberGroup.GetThaiWord()
	}
	return finalThaiWord
}
