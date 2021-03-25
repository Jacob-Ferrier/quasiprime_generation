package generate

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

type quasiprime struct {
	number       int
	moduloResult int
}

// Random change
type moduloData struct {
	quantity   int
	percentage float64
}

// QuasiprimeList structure
type QuasiprimeList struct {
	quasiprimes              map[int]quasiprime
	primes                   map[int]int
	nQuasiprime              int
	modulo                   int
	outFileName              string
	minIntegerChecked        int
	maxIntegerChecked        int
	numIntergersChecked      int
	minQuasiprime            int
	maxQuasiprime            int
	numQuasiprimes           int
	moduloDataList           map[int]moduloData
	pairedModuloDataList     map[int]map[int]moduloData
	quasiprimeGenerationTime time.Duration
}

func (quasiprimeList *QuasiprimeList) print() {
	fmt.Printf("Quasiprime List\n#######################\n")
	fmt.Printf("Primes used: %v\n", quasiprimeList.primes)
	fmt.Printf("Modulo: %v\n", quasiprimeList.modulo)
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
	quasiprimes := make(map[int]quasiprime)

	start := time.Now()
	for candidate := quasiprimeList.minIntegerChecked; candidate <= quasiprimeList.maxIntegerChecked; candidate++ {
		if isQuasiprime(candidate, quasiprimeList.nQuasiprime, quasiprimeList.primes) {
			quasiprimes[len(quasiprimes)] = quasiprime{candidate, candidate % quasiprimeList.modulo}
		}
	}
	quasiprimeList.quasiprimeGenerationTime = time.Since(start)

	quasiprimeList.quasiprimes = quasiprimes
	quasiprimeList.numQuasiprimes = len(quasiprimes)
	quasiprimeList.minQuasiprime = quasiprimes[0].number
	quasiprimeList.maxQuasiprime = quasiprimes[len(quasiprimes)-1].number

	for r := range quasiprimeList.moduloDataList {
		quantity := 0
		for _, quasiprime := range quasiprimes {
			if quasiprime.moduloResult == r {
				quantity++
			}
		}
		percentage := float64(quantity) / float64(quasiprimeList.numQuasiprimes)

		quasiprimeList.moduloDataList[r] = moduloData{quantity, percentage}
	}

	totalPairs := 0
	for i := 0; i < len(quasiprimeList.quasiprimes)-1; i++ {
		q, n := quasiprimeList.quasiprimes[i], quasiprimeList.quasiprimes[i+1]

		quasiprimeList.pairedModuloDataList[q.moduloResult][n.moduloResult] =
			moduloData{quasiprimeList.pairedModuloDataList[q.moduloResult][n.moduloResult].quantity + 1, 0.0}

		totalPairs++
	}

	for a := 0; a < len(quasiprimeList.pairedModuloDataList); a++ {
		for b := 0; b < len(quasiprimeList.pairedModuloDataList[a]); b++ {
			quasiprimeList.pairedModuloDataList[a][b] = moduloData{quasiprimeList.pairedModuloDataList[a][b].quantity,
				float64(quasiprimeList.pairedModuloDataList[a][b].quantity) / float64(totalPairs)}
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
	_, err = w.WriteString(fmt.Sprintf("#Modulo: %v\n", quasiprimeList.modulo))
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

	_, err = w.WriteString(fmt.Sprintf("\nModulo Result\tQuantitiy\tPercentage\n"))
	check(err)
	for i := 0; i < len(quasiprimeList.moduloDataList); i++ {
		_, err = w.WriteString(fmt.Sprintf("%v\t%v\t%v\n", i,
			quasiprimeList.moduloDataList[i].quantity, quasiprimeList.moduloDataList[i].percentage))
		check(err)
	}

	if complete {
		_, err = w.WriteString(fmt.Sprintf("\nPaired Modulo Result\tQuantity\tPercentage\n"))
		check(err)
		for a := 0; a < len(quasiprimeList.pairedModuloDataList); a++ {
			for b := 0; b < len(quasiprimeList.pairedModuloDataList[a]); b++ {
				_, err = w.WriteString(fmt.Sprintf("(%v,%v)\t%v\t%v\n", a, b,
					quasiprimeList.pairedModuloDataList[a][b].quantity, quasiprimeList.pairedModuloDataList[a][b].percentage))
				check(err)
			}
		}
	}

	_, err = w.WriteString(fmt.Sprintf("\nQuasiprime\tModulo Result\n"))
	check(err)
	for i := 0; i < len(quasiprimeList.quasiprimes); i++ {
		_, err = w.WriteString(fmt.Sprintf("%v\t%v\n", quasiprimeList.quasiprimes[i].number, quasiprimeList.quasiprimes[i].moduloResult))
		check(err)
	}

	w.Flush()
}

func worker(id int, quasiprimeList QuasiprimeList, writePrime bool, singleC chan map[int]moduloData, quasiprimeC chan map[int]quasiprime, timeC chan time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	quasiprimeList.generate()
	quasiprimeList.writeToFile(writePrime, false)
	fmt.Printf("Worker %d done\n", id)
	singleC <- quasiprimeList.moduloDataList
	quasiprimeC <- quasiprimeList.quasiprimes
	timeC <- quasiprimeList.quasiprimeGenerationTime
}

// Quasiprimes main generation function, generate quasiprimes up to maxNumberToGen with listSizeCaps
func Quasiprimes(v int, maxNumberToGen int, listSizeCap int, modulo int, nQuasiprime int, primeSourceFile string, outputDir string, writePrime bool) {
	// Initialize quasiprime lists
	vPrint(v, 0, "Initializing quasiprime lists\n####################\n")
	lists := makeIntitialLists(maxNumberToGen, listSizeCap, modulo, nQuasiprime, outputDir)
	vPrint(v, 0, "####################\nDone\n\n")

	// Make master prime list
	vPrint(v, 0, "Preloading prime list\n####################\n")
	masterPrimeList := makeMasterPrimeList(maxNumberToGen, primeSourceFile)
	vPrint(v, 0, "####################\nDone\n\n")

	// Make individual prime lists
	vPrint(v, 0, "Computing individual prime lists\n####################\n")
	for l := 0; l < len(lists); l++ {
		quasiprimeListToOperate := lists[l]
		quasiprimeListToOperate.getPrimeList(masterPrimeList)
		lists[l] = quasiprimeListToOperate
	}
	vPrint(v, 0, "####################\nDone\n\n")

	// Perform distributed computations
	vPrint(v, 0, "Distributing workloads to workers and deploying\n####################\n")
	var wg sync.WaitGroup

	singleC := make(chan map[int]moduloData, len(lists))
	quasiprimeC := make(chan map[int]quasiprime, len(lists))
	timeC := make(chan time.Duration, len(lists))
	for j := 0; j < len(lists); j++ {
		wg.Add(1)
		go worker(j, lists[j], writePrime, singleC, quasiprimeC, timeC, &wg)
	}

	wg.Wait()
	vPrint(v, 0, "All workers reported completion\n")
	close(singleC)
	close(quasiprimeC)
	close(timeC)
	vPrint(v, 0, "All channels closed\n")
	vPrint(v, 0, "####################\nDone\n\n")

	// Generate report
	vPrint(v, 0, "Generating final report\n####################\n")
	var completeQuasiprimeList QuasiprimeList

	completeModuloDataList := make(map[int]moduloData, modulo)
	for p := 0; p < modulo; p++ {
		completeModuloDataList[p] = moduloData{0, 0.0}
	}

	completeNumQuasiprimes := 0
	for u := range singleC {
		for r := 0; r < len(completeModuloDataList); r++ {
			completeModuloDataList[r] = moduloData{completeModuloDataList[r].quantity + u[r].quantity, 0.0}
			completeNumQuasiprimes += u[r].quantity
		}
	}

	var completeGenerationTime time.Duration
	for x := range timeC {
		completeGenerationTime += x
	}

	completeQuasiprimeList.modulo = modulo
	completeQuasiprimeList.nQuasiprime = nQuasiprime
	completeQuasiprimeList.outFileName = fmt.Sprintf("%s/%v-quasiprimes.modulo%v.complete_report.txt", outputDir, nQuasiprime, modulo)
	completeQuasiprimeList.minIntegerChecked = 0
	completeQuasiprimeList.maxIntegerChecked = maxNumberToGen
	completeQuasiprimeList.numIntergersChecked = maxNumberToGen + 1
	completeQuasiprimeList.numQuasiprimes = completeNumQuasiprimes
	completeQuasiprimeList.quasiprimeGenerationTime = completeGenerationTime

	for y := 0; y < len(completeModuloDataList); y++ {
		percentage := float64(completeModuloDataList[y].quantity) / float64(completeQuasiprimeList.numQuasiprimes)
		completeModuloDataList[y] = moduloData{completeModuloDataList[y].quantity, percentage}
	}

	completeQuasiprimeList.moduloDataList = completeModuloDataList

	coll := make(map[int]map[int]quasiprime, len(quasiprimeC))
	mins := make(map[int]int, len(quasiprimeC))
	order := make(map[int]int, len(quasiprimeC))
	i := 0
	for u := range quasiprimeC {
		coll[i] = u
		mins[i] = u[0].number
		order[i] = i
		i++
	}

	for j := 0; j < len(order)-1; j++ {
		swapped := false
		for k := 0; k < len(order)-1; k++ {
			if mins[order[k]] > mins[order[k+1]] {
				temp := order[k]
				order[k] = order[k+1]
				order[k+1] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	completeQuasiprimes := make(map[int]quasiprime, completeQuasiprimeList.numQuasiprimes)
	k := 0
	for i := 0; i < len(order); i++ {
		for j := 0; j < len(coll[order[i]]); j++ {
			completeQuasiprimes[k] = coll[order[i]][j]
			k++
		}
	}

	completeQuasiprimeList.quasiprimes = completeQuasiprimes
	completeQuasiprimeList.minQuasiprime = completeQuasiprimeList.quasiprimes[0].number
	completeQuasiprimeList.maxQuasiprime = completeQuasiprimeList.quasiprimes[completeQuasiprimeList.numQuasiprimes-1].number

	completePairedModuloDataList := make(map[int]map[int]moduloData, completeQuasiprimeList.modulo)
	for a := 0; a < modulo; a++ {
		completePairedModuloDataList[a] = make(map[int]moduloData, completeQuasiprimeList.modulo)
		for b := 0; b < modulo; b++ {
			completePairedModuloDataList[a][b] = moduloData{}
		}
	}
	completeQuasiprimeList.pairedModuloDataList = completePairedModuloDataList

	totalPairs := 0
	for i := 0; i < len(completeQuasiprimeList.quasiprimes)-1; i++ {
		q, n := completeQuasiprimeList.quasiprimes[i], completeQuasiprimeList.quasiprimes[i+1]

		completeQuasiprimeList.pairedModuloDataList[q.moduloResult][n.moduloResult] =
			moduloData{completeQuasiprimeList.pairedModuloDataList[q.moduloResult][n.moduloResult].quantity + 1, 0.0}

		totalPairs++
	}

	for a := 0; a < len(completeQuasiprimeList.pairedModuloDataList); a++ {
		for b := 0; b < len(completeQuasiprimeList.pairedModuloDataList[a]); b++ {
			completeQuasiprimeList.pairedModuloDataList[a][b] = moduloData{completeQuasiprimeList.pairedModuloDataList[a][b].quantity,
				float64(completeQuasiprimeList.pairedModuloDataList[a][b].quantity) / float64(totalPairs)}
		}
	}

	completeQuasiprimeList.writeToFile(false, true)
	vPrint(v, 0, "####################\nDone\n\n")

	fmt.Println("All done")
}
