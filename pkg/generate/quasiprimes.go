package generate

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type quasiprime struct {
	number       int
	moduloResult int
}

type moduloData struct {
	quantity   int
	percentage float64
}

// QuasiprimeList structure
type QuasiprimeList struct {
	quasiprimes         map[int]quasiprime
	primes              map[int]int
	modulo              int
	outFileName         string
	minIntegerChecked   int
	maxIntegerChecked   int
	numIntergersChecked int
	minQuasiprime       int
	maxQuasiprime       int
	numQuasiprimes      int
	moduloDataList      map[int]moduloData
}

func (quasiprimeList *QuasiprimeList) print() {
	fmt.Printf("Quasiprime List\n#######################\n")
	fmt.Printf("Primes used: %v\n", quasiprimeList.primes)
	fmt.Printf("Modulo: %v\n", quasiprimeList.modulo)
	fmt.Printf("Out File Name: %v\n", quasiprimeList.outFileName)
	fmt.Printf("Minimum Integer Checked: %v\n", quasiprimeList.minIntegerChecked)
	fmt.Printf("Maximum Integer Checked: %v\n", quasiprimeList.maxIntegerChecked)
	fmt.Printf("Number of Integers Checked: %v\n", quasiprimeList.numIntergersChecked)
	fmt.Printf("Minimum Quasiprime: %v\n", quasiprimeList.minQuasiprime)
	fmt.Printf("Maximum Quasiprime: %v\n", quasiprimeList.maxQuasiprime)
	fmt.Printf("Number of Quasiprimes Generated: %v\n", quasiprimeList.numQuasiprimes)
	fmt.Printf("Quasiprimes generated: %v\n", quasiprimeList.quasiprimes)
	fmt.Printf("Modulo Data: %v\n", quasiprimeList.moduloDataList)
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

	for candidate := quasiprimeList.minIntegerChecked; candidate <= quasiprimeList.maxIntegerChecked; candidate++ {
		if isQuasiprime(candidate, quasiprimeList.primes) {
			quasiprimes[len(quasiprimes)] = quasiprime{candidate, candidate % quasiprimeList.modulo}
		}
	}

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
}

func (quasiprimeList *QuasiprimeList) writeToFile() {
	f, err := os.Create(quasiprimeList.outFileName)
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(fmt.Sprintf("#Primes used: %v\n", quasiprimeList.primes))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("#Modulo: %v\n", quasiprimeList.modulo))
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
	_, err = w.WriteString(fmt.Sprintf("#Modulo Data: %v\n", quasiprimeList.moduloDataList))
	check(err)
	_, err = w.WriteString(fmt.Sprintf("Quasiprime\tModulo Result\n"))
	check(err)
	for i := 0; i < len(quasiprimeList.quasiprimes); i++ {
		_, err = w.WriteString(fmt.Sprintf("%v\t%v\n", quasiprimeList.quasiprimes[i].number, quasiprimeList.quasiprimes[i].moduloResult))
		check(err)
	}

	w.Flush()
}

func worker(id int, quasiprimeList QuasiprimeList, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	quasiprimeList.generate()
	quasiprimeList.writeToFile()
	fmt.Printf("Worker %d done\n", id)
}

// Quasiprimes main generation function, generate quasiprimes up to maxNumberToGen with listSizeCaps
func Quasiprimes(maxNumberToGen int, listSizeCap int, modulo int, primeSourceFile string, outputDir string) {
	numLists := int(math.Ceil(float64(maxNumberToGen) / float64(listSizeCap)))
	if numLists == maxNumberToGen/listSizeCap {
		numLists++
	}

	lists := make(map[int]QuasiprimeList, numLists)

	for i := 0; i < numLists; i++ {
		var quasiprimeList QuasiprimeList

		if ((i+1)*listSizeCap)-1 < maxNumberToGen {
			quasiprimeList.maxIntegerChecked = ((i + 1) * listSizeCap) - 1
		} else {
			quasiprimeList.maxIntegerChecked = maxNumberToGen
		}

		quasiprimeList.minIntegerChecked = i * listSizeCap
		quasiprimeList.numIntergersChecked = quasiprimeList.maxIntegerChecked - quasiprimeList.minIntegerChecked + 1
		quasiprimeList.outFileName = fmt.Sprintf("%s/quasiprimes.modulo%v.part.%016d.txt", outputDir, modulo, i)
		quasiprimeList.modulo = modulo

		moduloDataList := make(map[int]moduloData, modulo)
		for p := 0; p < modulo; p++ {
			moduloDataList[p] = moduloData{}
		}
		quasiprimeList.moduloDataList = moduloDataList

		lists[i] = quasiprimeList
	}

	// Preload prime list into memory
	fmt.Println("Preloading prime list")
	file, err := os.Open(primeSourceFile)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	masterPrimeList := make(map[int]int)
	maxPrimeNeeded := int(math.Ceil(float64(maxNumberToGen)/float64(2))) + 1
	stop := false
	for scanner.Scan() && !stop {
		line := scanner.Text()
		lineSlice := strings.Split(line, "\t")
		for _, k := range lineSlice {
			value, _ := strconv.Atoi(k)
			if value <= maxPrimeNeeded {
				masterPrimeList[len(masterPrimeList)] = value
			} else {
				stop = true
			}
		}
	}
	fmt.Println("Prime list preload done")

	fmt.Println("Computing individual prime lists")
	for l := 0; l < len(lists); l++ {
		quasiprimeListToOperate := lists[l]
		quasiprimeListToOperate.getPrimeList(masterPrimeList)
		lists[l] = quasiprimeListToOperate
	}
	fmt.Println("Done computing individual prime lists")

	var wg sync.WaitGroup

	for j := 0; j < len(lists); j++ {
		wg.Add(1)
		go worker(j, lists[j], &wg)
	}

	wg.Wait()

	fmt.Println("All workers reported completion")
}
