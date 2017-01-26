package summultiples

func MultipleSummer(multiples ...int) func(int) int {
	return func(limit int) int {
		sum := 0
		for i := 1; i < limit; i++ {
			for _, m := range multiples {
				if i%m == 0 {
					sum += i
					break
				}
			}
		}
		return sum
	}
}
