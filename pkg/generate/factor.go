package generate

import (
	"fmt"
	"math"
)

func isInteger(a float64) bool {
	return a == float64(int(a))
}

// nFactor function takes an candidate input and attempts to factor it into n prime factors recusively
// this function also utilizes a pregenerated prime list containing all of the necessary primes to validate primality
// of factors
// this function takes advantage of the sieve of Eratosthenes to dramatically cutdown on computation time
// in doing so, results can be very unpredictable for composites which do not have have n prime factors
// therefore this function is only to be used as a discrete test to see if a composite has n prime factors,
// it does not render any other accurate details other than a true/false evaluation
func nFactor(candidate float64, n int, startingPrimeIndex int, primes map[int]int, factorList []int) []int {
	fmt.Printf("Candidate: %v\n", candidate)

	if isPrime(candidate, primes) {
		factorList = append(factorList, int(candidate))
		fmt.Println("candidate is prime")
		return factorList
	}

	if n <= 0 {
		return factorList //prevent overflow safety step
	}

	temp := math.Pow(candidate, (1.0 / float64(n))) //sieve of Eratosthenes step, renders unpredictable results when factoring
	max := int(math.Floor(temp))
	if math.Pow(math.Ceil(temp), float64(n)) == candidate {
		max = int(math.Ceil(temp))
	}

	for i := startingPrimeIndex; primes[i] <= max; i++ {
		result := candidate / float64(primes[i])
		if isInteger(result) {
			factorList = append(factorList, primes[i])
			factorList = nFactor(result, n-1, 0, primes, factorList)
			break
		}

	}

	return factorList
}

func Factor(candidate int, n int, primes map[int]int) []int {
	var factorList []int
	factorList = nFactor(float64(candidate), n, 0, primes, factorList)

	return factorList

}
