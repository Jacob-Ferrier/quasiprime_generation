package main

import "github.com/Jacob-Ferrier/quasiprime_generation/pkg/generate"

func main() {
	generate.Quasiprimes(67, 14, 4, []int{1, 3},
		"/home/jacobferrier/Desktop/prime_computations_test/data/in/primes/combined/2T_combined.txt",
		"/home/jacobferrier/Desktop/quasiprime_generation_test")
}
