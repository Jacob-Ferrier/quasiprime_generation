package generate

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

type moduloData struct {
	quantity   int
	percentage float64
}

// QuasiprimeList structure
type QuasiprimeList struct {
	quasiprimes              map[int]int
	primes                   map[int]int
	nQuasiprime              int
	moduloMax                int
	outFileName              string
	minIntegerChecked        int
	maxIntegerChecked        int
	numIntergersChecked      int
	minQuasiprime            int
	maxQuasiprime            int
	numQuasiprimes           int
	moduloDataList           map[int]map[int]moduloData
	pairedModuloDataList     map[int]map[int]map[int]moduloData
	quasiprimeGenerationTime time.Duration
}

func (quasiprimeList *QuasiprimeList) print() {
	fmt.Printf("Quasiprime List\n#######################\n")
	fmt.Printf("Primes used: %v\n", quasiprimeList.primes)
	fmt.Printf("Modulating up to: %v\n", quasiprimeList.moduloMax)
	fmt.Printf("Number of Prime Factors in Quasiprimes: %v\n", quasiprimeList.nQuasiprime)
	fmt.Printf("Out File Name: %v\n", quasiprimeList.outFileName)
	fmt.Printf("Minimum Integer Checked: %v\n", quasiprimeList.minIntegerChecked)
	fmt.Printf("Maximum Integer Checked: %v\n", quasiprimeList.maxIntegerChecked)
	fmt.Printf("Number of Integers Checked: %v\n", quasiprimeList.numIntergersChecked)
	fmt.Printf("Minimum Quasiprime: %v\n", quasiprimeList.minQuasiprime)
	fmt.Printf("Maximum Quasiprime: %v\n", quasiprimeList.maxQuasiprime)
	fmt.Printf("Number of Quasiprimes Generated: %v\n", quasiprimeList.numQuasiprimes)
	fmt.Printf("Quasiprimes generated: %v\n", quasiprimeList.quasiprimes)
	fmt.Printf("Quasiprime Generation Time: %v\n", quasiprimeList.quasiprimeGenerationTime)
	fmt.Printf("Modulo Data: %v\n", quasiprimeList.moduloDataList)
	fmt.Printf("Paired Modulo Data: %v\n", quasiprimeList.pairedModuloDataList)
	fmt.Printf("\n")
}

func (quasiprimeList *QuasiprimeList) getPrimeList(masterPrimeList map[int]int) {
	requiredMax := int(math.Ceil(float64(quasiprimeList.maxIntegerChecked)/float64(2))) + 1
	primes := make(map[int]int)
	for i := 0; i < len(masterPrimeList); i++ {
		if masterPrimeList[i] <= requiredMax {
			primes[len(primes)] = masterPrimeList[i]
		} else {
			break
		}
	}

	quasiprimeList.primes = primes
}

func (quasiprimeList *QuasiprimeList) generate() {
	quasiprimes := make(map[int]int)

	start := time.Now()
	for candidate := quasiprimeList.minIntegerChecked; candidate <= quasiprimeList.maxIntegerChecked; candidate++ {
		if isQuasiprime(candidate, quasiprimeList.nQuasiprime, quasiprimeList.primes) {
			quasiprimes[len(quasiprimes)] = candidate
		}
	}
	quasiprimeList.quasiprimeGenerationTime = time.Since(start)

	quasiprimeList.quasiprimes = quasiprimes
	quasiprimeList.numQuasiprimes = len(quasiprimes)
	quasiprimeList.minQuasiprime = quasiprimes[0]
	quasiprimeList.maxQuasiprime = quasiprimes[len(quasiprimes)-1]

	for m := range quasiprimeList.moduloDataList {
		for r := range quasiprimeList.moduloDataList[m] {
			quantity := 0
			for _, quasiprime := range quasiprimes {
				if quasiprime%m == r {
					quantity++
				}
			}
			percentage := float64(quantity) / float64(quasiprimeList.numQuasiprimes)

			quasiprimeList.moduloDataList[m][r] = moduloData{quantity, percentage}
		}
	}

	for m := range quasiprimeList.pairedModuloDataList {
		totalPairs := 0
		for i := 0; i < len(quasiprimeList.quasiprimes)-1; i++ {
			q, n := quasiprimeList.quasiprimes[i], quasiprimeList.quasiprimes[i+1]

			quasiprimeList.pairedModuloDataList[m][q%m][n%m] =
				moduloData{quasiprimeList.pairedModuloDataList[m][q%m][n%m].quantity + 1, 0.0}

			totalPairs++
		}

		for a := 0; a < len(quasiprimeList.pairedModuloDataList[m]); a++ {
			for b := 0; b < len(quasiprimeList.pairedModuloDataList[m][a]); b++ {
				quasiprimeList.pairedModuloDataList[m][a][b] = moduloData{quasiprimeList.pairedModuloDataList[m][a][b].quantity,
					float64(quasiprimeList.pairedModuloDataList[m][a][b].quantity) / float64(totalPairs)}
			}
		}
	}
}

func (quasiprimeList *QuasiprimeList) writeToFile(writePrime bool, complete bool) {
	f, err := os.Create(quasiprimeList.outFileName)
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)

	if writePrime {
		_, err = w.WriteString(fmt.Sprintf("#Primes used: %v\n", quasiprimeList.primes))
		check(err)
	}

	_, err = w.WriteString(fmt.Sprintf("#Modulating up to: %v\n", quasiprimeList.moduloMax))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Number of Prime Factors in Quasiprimes: %v\n", quasiprimeList.nQuasiprime))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Out File Name: %v\n", quasiprimeList.outFileName))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Minimum Integer Checked: %v\n", quasiprimeList.minIntegerChecked))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Maximum Integer Checked: %v\n", quasiprimeList.maxIntegerChecked))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Number of Integers Checked: %v\n", quasiprimeList.numIntergersChecked))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Minimum Quasiprime: %v\n", quasiprimeList.minQuasiprime))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Maximum Quasiprime: %v\n", quasiprimeList.maxQuasiprime))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Number of Quasiprimes Generated: %v\n", quasiprimeList.numQuasiprimes))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Quasiprime Generation Time: %v\n", quasiprimeList.quasiprimeGenerationTime))
	check(err)

	for m := 2; m < len(quasiprimeList.moduloDataList)+2; m++ {
		_, err = w.WriteString(fmt.Sprintf("\nModulo(%v) Result\tQuantitiy\tPercentage\n", m))
		check(err)
		for i := 0; i < len(quasiprimeList.moduloDataList[m]); i++ {
			_, err = w.WriteString(fmt.Sprintf("%v\t%v\t%v\n", i,
				quasiprimeList.moduloDataList[m][i].quantity, quasiprimeList.moduloDataList[m][i].percentage))
			check(err)
		}
	}

	if complete {
		for m := 2; m < len(quasiprimeList.pairedModuloDataList)+2; m++ {
			_, err = w.WriteString(fmt.Sprintf("\nPaired Modulo(%v) Result\tQuantity\tPercentage\n", m))
			check(err)
			for a := 0; a < m; a++ {
				for b := 0; b < m; b++ {
					_, err = w.WriteString(fmt.Sprintf("(%v,%v)\t%v\t%v\n", a, b,
						quasiprimeList.pairedModuloDataList[m][a][b].quantity, quasiprimeList.pairedModuloDataList[m][a][b].percentage))
					check(err)
				}
			}
		}
	}

	_, err = w.WriteString(fmt.Sprintf("\nQuasiprime"))
	check(err)

	for m := 2; m <= quasiprimeList.moduloMax; m++ {
		_, err = w.WriteString(fmt.Sprintf("\t(%v)", m))
		check(err)
	}

	for i := 0; i < len(quasiprimeList.quasiprimes); i++ {
		quasiprime := quasiprimeList.quasiprimes[i]
		_, err = w.WriteString(fmt.Sprintf("\n%v", quasiprime))
		check(err)

		for m := 2; m <= quasiprimeList.moduloMax; m++ {
			_, err = w.WriteString(fmt.Sprintf("\t%v", quasiprime%m))
			check(err)
		}
	}

	w.Flush()
}

func worker(id int, quasiprimeList QuasiprimeList, writePrime bool, quasiprimeListChannel chan QuasiprimeList, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	quasiprimeList.generate()
	quasiprimeList.writeToFile(writePrime, true)
	fmt.Printf("Worker %d done\n", id)
	quasiprimeListChannel <- quasiprimeList
}

// Quasiprimes main generation function, generate quasiprimes up to maxNumberToGen with listSizeCaps
func Quasiprimes(v int, maxNumberToGen int, listSizeCap int, moduloMax int, nQuasiprime int, primeSourceFile string, outputDir string, writePrime bool) {
	// Initialize quasiprime lists
	vPrint(v, 0, "Initializing quasiprime lists\n####################\n")
	lists := makeIntitialLists(maxNumberToGen, listSizeCap, moduloMax, nQuasiprime, outputDir)
	vPrint(v, 0, "####################\nDone\n\n")
	/////////////////////////////////////

	// Make master prime list
	vPrint(v, 0, "Preloading prime list\n####################\n")
	masterPrimeList := makeMasterPrimeList(maxNumberToGen, primeSourceFile)
	vPrint(v, 0, "####################\nDone\n\n")
	/////////////////////////////////////

	// Make individual prime lists
	vPrint(v, 0, "Computing individual prime lists\n####################\n")
	for l := 0; l < len(lists); l++ {
		quasiprimeListToOperate := lists[l]
		quasiprimeListToOperate.getPrimeList(masterPrimeList)
		lists[l] = quasiprimeListToOperate
	}
	vPrint(v, 0, "####################\nDone\n\n")
	/////////////////////////////////////

	// Perform distributed computations
	vPrint(v, 0, "Distributing workloads to workers and deploying\n####################\n")

	var wg sync.WaitGroup
	quasiprimeListChannel := make(chan QuasiprimeList, len(lists))
	for j := 0; j < len(lists); j++ {
		wg.Add(1)
		go worker(j, lists[j], writePrime, quasiprimeListChannel, &wg)
	}

	wg.Wait()
	vPrint(v, 0, "All workers reported completion\n")
	close(quasiprimeListChannel)
	vPrint(v, 0, "All channels closed\n")
	vPrint(v, 0, "####################\nDone\n\n")
	/////////////////////////////////////

	// Generate report
	vPrint(v, 0, "Generating final report\n####################\n")
	report(maxNumberToGen, listSizeCap, moduloMax, nQuasiprime, outputDir, quasiprimeListChannel)
	vPrint(v, 0, "####################\nDone\n\n")
	/////////////////////////////////////

	fmt.Println("All done")
}
