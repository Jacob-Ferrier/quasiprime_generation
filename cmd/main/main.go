package main

import (
	"flag"

	"github.com/Jacob-Ferrier/quasiprime_generation/pkg/generate"
)

func main() {
	// flag defitions
	var vFlag = flag.Int("v", 0, `verbosity level,
		(-1) no messages,
		(0) normal messsages,
		(1) normal messages with more information,
		(2) normal messages with detailed information,
		(3) normal messages with full details,
		(4) dump to output (not recommended),
		(5) debugging`)
	var maxNumberToGenFlag = flag.Int("maxNumberToGen", 100, "maximum number (not quantity) to generate up to")
	var listSizeCapFlag = flag.Int("listSizeCap", 101, "maximum size of a quasiprime list, set listSizeCap to maxNumberToGen+1 for a single list")
	var moduloMaxFlag = flag.Int("moduloMax", 4, "starting with 2, number to modulate up to")
	var nQuasiprimeFlag = flag.Int("nQuasiprime", 2, "number of prime factors in quasiprimes")
	var primeSourceFileFlag = flag.String("primeSourceFile", "", "full path to the prime source file")
	var outputDirFlag = flag.String("outputDir", "", "full path to output dir")
	var writePrimeFlag = flag.Bool("writePrime", false, "write primes used during computation to output files")

	// parse command line flags
	flag.Parse()

	// main runtime
	generate.Quasiprimes(*vFlag, *maxNumberToGenFlag, *listSizeCapFlag, *moduloMaxFlag, *nQuasiprimeFlag, *primeSourceFileFlag, *outputDirFlag, *writePrimeFlag)
}
