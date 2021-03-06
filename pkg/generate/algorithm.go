package generate

import (
	"math"
)

func isQuasiprime(candidate int, nQuasiprime int, primes map[int]int) bool {
	primeFactors := Factor(candidate, nQuasiprime, primes)

	temp := 1
	for _, p := range primeFactors {
		temp *= p
	}

	return len(primeFactors) == nQuasiprime && temp == candidate
}

func isPrime(result float64, primes map[int]int) bool {
	if float64(result) == math.Floor(result) {
		if contains(int(result), primes) {
			return true
		}
	}
	return false
}

func contains(value int, primes map[int]int) bool {
	for i := 0; i < len(primes); i++ {
		if value == primes[i] {
			return true
		}
	}

	return false
}
