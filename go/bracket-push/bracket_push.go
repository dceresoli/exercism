package brackets

const testVersion = 4

func Bracket(s string) (bool, error) {
	bra := "([{"
	ket := ")]}"
	list := make([]byte, 0)

	for _, c := range s {
		// check for bra
		for _, b := range bra {
			if c == b {
				list = append(list, byte(c))
			}
		}

		// check for ket
		for i, k := range ket {
			if c == k {
				if len(list) == 0 {
					return false, nil
				}
				if list[len(list)-1] != bra[i] {
					return false, nil
				}
				list = list[0 : len(list)-1]
			}
		}

	}

	if len(list) > 0 {
		return false, nil
	}

	return true, nil
}
