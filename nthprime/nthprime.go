package nthprime

import "math"

func NthPrime(n int) int {
	primes := []int{}
	for num := 2; ; num++ {
		if IsPrime(num) {
			primes = append(primes, num)
			if len(primes)-1 == n {
				return num
			}
		}
	}
}

func IsPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	upper := int(math.Sqrt(float64(n)))
	for m := 2; m <= upper; m++ {
		if n%m == 0 {
			return false
		}
	}
	return true
}
