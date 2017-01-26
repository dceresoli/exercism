package sieve

func Sieve(limit int) []int {
	// initialize list and assume every number is prime
	isprime := make([]bool, limit+1)
	for i := 0; i <= limit; i++ {
		isprime[i] = true
	}

	// mark prime numbers
	isprime[0] = true
	isprime[1] = true

	// sieve
	for i := 2; i < limit; i++ {
		if !isprime[i] {
			continue
		}
		isprime[i] = true
		for j := i * 2; j <= limit; j += i {
			isprime[j] = false
		}
	}

	// return list of primes
	primes := make([]int, 0)
	for i := 2; i <= limit; i++ {
		if isprime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
