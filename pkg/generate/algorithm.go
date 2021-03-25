package generate

import (
	"math"
)

func isQuasiprime(candidate int, primes map[int]int) bool {
	max := int(math.Floor(math.Sqrt(float64(candidate)))) // by the Sieve of Eratosthenes

	for j := 0; primes[j] <= max; j++ {
		result := float64(candidate) / float64(primes[j])
		if isPrime(result, primes) {
			return true
		}
	}

	return false
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
