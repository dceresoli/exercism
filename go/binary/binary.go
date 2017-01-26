package binary

import "errors"

const testVersion = 2

func ParseBinary(s string) (int, error) {
	bin := 0
	for _, c := range s {
		if c == '0' {
			bin = bin << 1
		} else if c == '1' {
			bin = bin<<1 + 1
		} else {
			return 0, errors.New("wrong binary string")
		}
	}
	return bin, nil
}
