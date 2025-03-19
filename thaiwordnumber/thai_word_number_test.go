package thaiwordnumber

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGetThaiWordNumberOf_wholeNumber(t *testing.T) {
	thaiWordNumber := NewThaiWordNumber()

	t.Run("Test whole number 1", func(t *testing.T) {
		v, err := decimal.NewFromString("1")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งบาทถ้วน", thaiWord)
	})
	t.Run("Test whole number 2", func(t *testing.T) {
		v, err := decimal.NewFromString("2")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สองบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 3", func(t *testing.T) {
		v, err := decimal.NewFromString("3")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สามบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 4", func(t *testing.T) {
		v, err := decimal.NewFromString("4")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สี่บาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 5", func(t *testing.T) {
		v, err := decimal.NewFromString("5")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ห้าบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 6", func(t *testing.T) {
		v, err := decimal.NewFromString("6")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หกบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 7", func(t *testing.T) {
		v, err := decimal.NewFromString("7")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "เจ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 8", func(t *testing.T) {
		v, err := decimal.NewFromString("8")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "แปดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 9", func(t *testing.T) {
		v, err := decimal.NewFromString("9")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "เก้าบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 10", func(t *testing.T) {
		v, err := decimal.NewFromString("10")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 11", func(t *testing.T) {
		v, err := decimal.NewFromString("11")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 21", func(t *testing.T) {
		v, err := decimal.NewFromString("21")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ยี่สิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 100", func(t *testing.T) {
		v, err := decimal.NewFromString("100")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยบาทถ้วน", thaiWord)
	})
	t.Run("Test whole number 101", func(t *testing.T) {
		v, err := decimal.NewFromString("101")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 110", func(t *testing.T) {
		v, err := decimal.NewFromString("110")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยสิบบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 111", func(t *testing.T) {
		v, err := decimal.NewFromString("111")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยสิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 1,001", func(t *testing.T) {
		v, err := decimal.NewFromString("1001")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งพันเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 1,111", func(t *testing.T) {
		v, err := decimal.NewFromString("1111")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งพันหนึ่งร้อยสิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 12,121", func(t *testing.T) {
		v, err := decimal.NewFromString("12121")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งหมื่นสองพันหนึ่งร้อยยี่สิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 10,121", func(t *testing.T) {
		v, err := decimal.NewFromString("10121")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งหมื่นหนึ่งร้อยยี่สิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 121,451", func(t *testing.T) {
		v, err := decimal.NewFromString("121451")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งแสนสองหมื่นหนึ่งพันสี่ร้อยห้าสิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 1,234,511", func(t *testing.T) {
		v, err := decimal.NewFromString("1234511")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยสิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 12,315,111", func(t *testing.T) {
		v, err := decimal.NewFromString("12315111")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบสองล้านสามแสนหนึ่งหมื่นห้าพันหนึ่งร้อยสิบเอ็ดบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 123,056,789", func(t *testing.T) {
		v, err := decimal.NewFromString("123056789")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยยี่สิบสามล้านห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 1,234,567,890", func(t *testing.T) {
		v, err := decimal.NewFromString("1234567890")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 12,234,567,890", func(t *testing.T) {
		v, err := decimal.NewFromString("12234567890")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งหมื่นสองพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 123,456,789,012", func(t *testing.T) {
		v, err := decimal.NewFromString("123456789012")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านเจ็ดแสนแปดหมื่นเก้าพันสิบสองบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 2,123,456,789,012", func(t *testing.T) {
		v, err := decimal.NewFromString("2123456789012")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สองล้านหนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านเจ็ดแสนแปดหมื่นเก้าพันสิบสองบาทถ้วน", thaiWord)
	})
}

func TestGetThaiWordNumberOf_wholeNumber_with_decimal(t *testing.T) {
	thaiWordNumber := NewThaiWordNumber()

	t.Run("Test whole number 0.01", func(t *testing.T) {
		v, err := decimal.NewFromString("0.01")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งสตางค์", thaiWord)
	})

	t.Run("Test whole number 0.11", func(t *testing.T) {
		v, err := decimal.NewFromString("0.11")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบเอ็ดสตางค์", thaiWord)
	})

	t.Run("Test whole number 0.55", func(t *testing.T) {
		v, err := decimal.NewFromString("0.55")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ห้าสิบห้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 0.5", func(t *testing.T) {
		v, err := decimal.NewFromString("0.5")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ห้าสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 0.50", func(t *testing.T) {
		v, err := decimal.NewFromString("0.50")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ห้าสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 1.71", func(t *testing.T) {
		v, err := decimal.NewFromString("1.71")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งบาทเจ็ดสิบเอ็ดสตางค์", thaiWord)
	})
	t.Run("Test whole number 2.01", func(t *testing.T) {
		v, err := decimal.NewFromString("2.01")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สองบาทหนึ่งสตางค์", thaiWord)
	})

	t.Run("Test whole number 3.11", func(t *testing.T) {
		v, err := decimal.NewFromString("3.11")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สามบาทสิบเอ็ดสตางค์", thaiWord)
	})

	t.Run("Test whole number 4.21", func(t *testing.T) {
		v, err := decimal.NewFromString("4.21")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สี่บาทยี่สิบเอ็ดสตางค์", thaiWord)
	})

	t.Run("Test whole number 5.00", func(t *testing.T) {
		v, err := decimal.NewFromString("5.00")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ห้าบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 6.0", func(t *testing.T) {
		v, err := decimal.NewFromString("6.0")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หกบาทถ้วน", thaiWord)
	})

	t.Run("Test whole number 7.32", func(t *testing.T) {
		v, err := decimal.NewFromString("7.32")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "เจ็ดบาทสามสิบสองสตางค์", thaiWord)
	})

	t.Run("Test whole number 8.43", func(t *testing.T) {
		v, err := decimal.NewFromString("8.43")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "แปดบาทสี่สิบสามสตางค์", thaiWord)
	})

	t.Run("Test whole number 9.10", func(t *testing.T) {
		v, err := decimal.NewFromString("9.10")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "เก้าบาทสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 10.50", func(t *testing.T) {
		v, err := decimal.NewFromString("10.50")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบบาทห้าสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 11.5", func(t *testing.T) {
		v, err := decimal.NewFromString("11.5")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบเอ็ดบาทห้าสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 21.49", func(t *testing.T) {
		v, err := decimal.NewFromString("21.49")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "ยี่สิบเอ็ดบาทสี่สิบเก้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 100.02", func(t *testing.T) {
		v, err := decimal.NewFromString("100.02")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยบาทสองสตางค์", thaiWord)
	})
	t.Run("Test whole number 101.61", func(t *testing.T) {
		v, err := decimal.NewFromString("101.61")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยเอ็ดบาทหกสิบเอ็ดสตางค์", thaiWord)
	})

	t.Run("Test whole number 110.99", func(t *testing.T) {
		v, err := decimal.NewFromString("110.99")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยสิบบาทเก้าสิบเก้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 111.29", func(t *testing.T) {
		v, err := decimal.NewFromString("111.29")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยสิบเอ็ดบาทยี่สิบเก้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 1,001.01", func(t *testing.T) {
		v, err := decimal.NewFromString("1001.01")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งพันเอ็ดบาทหนึ่งสตางค์", thaiWord)
	})

	t.Run("Test whole number 1,111.11", func(t *testing.T) {
		v, err := decimal.NewFromString("1111.11")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งพันหนึ่งร้อยสิบเอ็ดบาทสิบเอ็ดสตางค์", thaiWord)
	})

	t.Run("Test whole number 12,121.89", func(t *testing.T) {
		v, err := decimal.NewFromString("12121.89")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งหมื่นสองพันหนึ่งร้อยยี่สิบเอ็ดบาทแปดสิบเก้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 10,121.80", func(t *testing.T) {
		v, err := decimal.NewFromString("10121.80")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งหมื่นหนึ่งร้อยยี่สิบเอ็ดบาทแปดสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 121,451.55", func(t *testing.T) {
		v, err := decimal.NewFromString("121451.55")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งแสนสองหมื่นหนึ่งพันสี่ร้อยห้าสิบเอ็ดบาทห้าสิบห้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 1,234,511.99", func(t *testing.T) {
		v, err := decimal.NewFromString("1234511.99")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยสิบเอ็ดบาทเก้าสิบเก้าสตางค์", thaiWord)
	})

	t.Run("Test whole number 12,315,111.22", func(t *testing.T) {
		v, err := decimal.NewFromString("12315111.22")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สิบสองล้านสามแสนหนึ่งหมื่นห้าพันหนึ่งร้อยสิบเอ็ดบาทยี่สิบสองสตางค์", thaiWord)
	})

	t.Run("Test whole number 123,056,789.71", func(t *testing.T) {
		v, err := decimal.NewFromString("123056789.71")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งร้อยยี่สิบสามล้านห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเจ็ดสิบเอ็ดสตางค์", thaiWord)
	})

	t.Run("Test whole number 1,234,567,890.78", func(t *testing.T) {
		v, err := decimal.NewFromString("1234567890.78")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทเจ็ดสิบแปดสตางค์", thaiWord)
	})

	t.Run("Test whole number 12,234,567,890.90", func(t *testing.T) {
		v, err := decimal.NewFromString("12234567890.90")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งหมื่นสองพันสองร้อยสามสิบสี่ล้านห้าแสนหกหมื่นเจ็ดพันแปดร้อยเก้าสิบบาทเก้าสิบสตางค์", thaiWord)
	})

	t.Run("Test whole number 123,456,789,012.12", func(t *testing.T) {
		v, err := decimal.NewFromString("123456789012.12")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านเจ็ดแสนแปดหมื่นเก้าพันสิบสองบาทสิบสองสตางค์", thaiWord)
	})

	t.Run("Test whole number 2,123,456,789,012.12", func(t *testing.T) {
		v, err := decimal.NewFromString("2123456789012.12")
		if err != nil {
			panic(err)
		}
		thaiWord := thaiWordNumber.GetThaiWordNumberOf(v)
		assert.Equal(t, "สองล้านหนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกล้านเจ็ดแสนแปดหมื่นเก้าพันสิบสองบาทสิบสองสตางค์", thaiWord)
	})

}
