package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

const testVersion = 1

func Number(num string) (out string, err error) {
	out = ""
	for _, c := range num {
		if unicode.IsDigit(c) {
			out += string(c)
		}
	}

	if out[0] == '1' && len(out) == 11 {
		out = out[1:]
	}
	if len(out) != 10 {
		return out, errors.New("Invalid phone number")
	}
	return out, nil
}

func AreaCode(num string) (area string, err error) {
	num, err = Number(num)
	if err != nil {
		return "", err
	}
	area = num[0:3]
	return area, nil
}

func Format(num string) (fnum string, err error) {
	num, err = Number(num)
	if err != nil {
		return "", err
	}
	fnum = fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:])
	return fnum, nil
}
