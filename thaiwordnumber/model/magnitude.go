package model

import "belerionbe/thaiwordnumber/constant"

type Magnitude struct {
	unitMap map[string]string
}

func NewMagnitude() *Magnitude {
	u := &Magnitude{
		unitMap: map[string]string{},
	}
	u.unitMap[constant.OneHundredThousand] = constant.OneHundredThousandReadAs
	u.unitMap[constant.TenThousand] = constant.TenThousandReadAS
	u.unitMap[constant.OneThousand] = constant.OneThousandReadAs
	u.unitMap[constant.OneHundred] = constant.OneHundredReadAs
	u.unitMap[constant.Ten] = constant.TenReadAs
	u.unitMap[constant.One] = ""
	return u
}

// Get the unit for the given value like แสน หมื่น พัน ร้อย สิบ
func (m *Magnitude) Get(value string) string {
	return m.unitMap[value]
}
