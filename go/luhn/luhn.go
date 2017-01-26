package luhn

import (
	"strconv"
)

// compute luhn cheksum
func calcChecksum(luhn string) (int, bool) {
	digits := make([]int, 0)
	for _, c := range luhn {
		i, err := strconv.Atoi(string(c))
		if err == nil {
			digits = append(digits, i)
		}
	}
	if len(digits) == 0 {
		return 0, false
	}

	// double every two digits
	for i := len(digits) - 2; i >= 0; i -= 2 {
		digits[i] = 2 * digits[i]
		if digits[i] > 9 {
			digits[i] -= 9
		}
	}

	// compute checksum
	checksum := 0
	for i := 0; i < len(digits); i++ {
		checksum += digits[i]
	}
	return checksum, true
}

// check if valid
func Valid(luhn string) bool {
	checksum, ok := calcChecksum(luhn)
	if !ok {
		return false
	}
	return checksum%10 == 0
}

// add control digit
func AddCheck(raw string) string {
	checksum, ok := calcChecksum(raw + "0")
	if !ok {
		return ""
	}
	diff := (10 - checksum%10) % 10
	return raw + strconv.Itoa(diff)
}
