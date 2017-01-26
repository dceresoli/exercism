package romannumerals

import "errors"

const testVersion = 3

func appendN(s string, c string, n int) string {
	for i := 0; i < n; i++ {
		s += c
	}
	return s
}

func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3999 {
		return "", errors.New("out of range")
	}

	// extract units, tens, hundreds, ...
	units := n % 10
	decis := (n / 10) % 10
	hunds := (n / 100) % 10
	mills := n / 1000

	res := ""

	// thousands
	res = appendN(res, "M", mills)

	// hundreds
	if hunds == 9 {
		res += "CM"
	} else if hunds >= 5 && hunds <= 8 {
		res += "D"
		res = appendN(res, "C", hunds-5)
	} else if hunds == 4 {
		res += "CD"
	} else if hunds <= 3 {
		res = appendN(res, "C", hunds)
	}

	// tens
	if decis == 9 {
		res += "XC"
	} else if decis >= 5 && decis <= 8 {
		res += "L"
		res = appendN(res, "X", decis-5)
	} else if decis == 4 {
		res += "XL"
	} else if decis <= 3 {
		res = appendN(res, "X", decis)
	}

	// units
	if units == 9 {
		res += "IX"
	} else if units >= 5 && units <= 8 {
		res += "V"
		res = appendN(res, "I", units-5)
	} else if units == 4 {
		res += "IV"
	} else if units <= 3 {
		res = appendN(res, "I", units)
	}
	return res, nil
}
