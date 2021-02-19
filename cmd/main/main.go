package main

import (
	"flag"

	"github.com/Jacob-Ferrier/quasiprime_generation/pkg/generate"
)

func main() {
	// flag defitions
	var maxNumberToGenFlag = flag.Int("maxNumberToGen", 100, "maximum number (not quantity) to generate up to")
	var listSizeCapFlag = flag.Int("listSizeCap", 101, "maximum size of a quasiprime list, set listSizeCap to maxNumberToGen+1 for a single list")
	var moduloFlag = flag.Int("modulo", 4, "number to modulate by")
	var primeSourceFileFlag = flag.String("primeSourceFile", "", "full path to the prime source file")
	var outputDirFlag = flag.String("outputDir", "", "full path to output dir")

	// parse command line flags
	flag.Parse()

	// main runtime
	generate.Quasiprimes(*maxNumberToGenFlag, *listSizeCapFlag, *moduloFlag, *primeSourceFileFlag, *outputDirFlag)
}
