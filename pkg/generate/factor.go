package generate

import (
	"fmt"
	"math"
)

func isInteger(a float64) bool {
	return a == float64(int(a))
}

func nFactor(candidate float64, n int, primes map[int]int, factorList []int) []int {
	if isPrime(candidate, primes) {
		factorList = append(factorList, int(candidate))
		return factorList
	}
	max := int(math.Floor(math.Pow(candidate, (1.0 / float64(n)))))

	for i := 0; primes[i] <= max; i++ {
		result := candidate / float64(primes[i])
		if isInteger(result) {
			factorList = append(factorList, primes[i])
			factorList = nFactor(result, n-1, primes, factorList)
		}
	}

	return factorList
}

func Factor(candidate int, n int, primes map[int]int) {
	var factorList []int
	factorList = nFactor(float64(candidate), n, primes, factorList)

	fmt.Println(factorList)

}
