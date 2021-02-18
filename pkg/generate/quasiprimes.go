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

// QuasiprimeList structure
type QuasiprimeList struct {
	quasiprimes         map[int]quasiprime
	primes              map[int]int
	modulo              int
	moduloResultOptions []int
	outFileName         string
	minIntegerChecked   int
	maxIntegerChecked   int
	numIntergersChecked int
	minQuasiprime       int
	maxQuasiprime       int
	numQuasiprimes      int
}

func (quasiprimeList *QuasiprimeList) print() {
	fmt.Printf("Quasiprime List\n#######################\n")
	fmt.Printf("Primes used: %v\n", quasiprimeList.primes)
	fmt.Printf("Modulo: %v\n", quasiprimeList.modulo)
	fmt.Printf("Modulo Result Options: %v\n", quasiprimeList.moduloResultOptions)
	fmt.Printf("Out File Name: %v\n", quasiprimeList.outFileName)
	fmt.Printf("Minimum Integer Checked: %v\n", quasiprimeList.minIntegerChecked)
	fmt.Printf("Maximum Integer Checked: %v\n", quasiprimeList.maxIntegerChecked)
	fmt.Printf("Number of Integers Checked: %v\n", quasiprimeList.numIntergersChecked)
	fmt.Printf("Minimum Quasiprime: %v\n", quasiprimeList.minQuasiprime)
	fmt.Printf("Maximum Quasiprime: %v\n", quasiprimeList.maxQuasiprime)
	fmt.Printf("Number of Quasiprimes Generated: %v\n", quasiprimeList.numQuasiprimes)
	fmt.Printf("Quasiprimes generated: %v\n", quasiprimeList.quasiprimes)
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
	//quasiprimes := make(map[int]quasiprime)

	for candidate := quasiprimeList.minIntegerChecked; candidate <= quasiprimeList.maxIntegerChecked; candidate++ {
		fmt.Println(candidate)
	}

}

func (quasiprimeList *QuasiprimeList) writeToFile() {
}

func worker(id int, quasiprimeList QuasiprimeList, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	quasiprimeList.generate()
	quasiprimeList.writeToFile()
	fmt.Printf("Worker %d done\n", id)
}

// Quasiprimes main generation function, generate quasiprimes up to maxNumberToGen with listSizeCaps
func Quasiprimes(maxNumberToGen int, listSizeCap int, modulo int, moduloResultOptions []int, primeSourceFile string, outputDir string) {
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
		quasiprimeList.outFileName = fmt.Sprintf("%s/quasiprimes.part.%016d.txt", outputDir, i)
		quasiprimeList.modulo = modulo
		quasiprimeList.moduloResultOptions = moduloResultOptions

		lists[i] = quasiprimeList
	}

	// Preload prime list into memory
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

	for l := 0; l < len(lists); l++ {
		quasiprimeListToOperate := lists[l]
		quasiprimeListToOperate.getPrimeList(masterPrimeList)
		quasiprimeListToOperate.generate()
	}

	//var wg sync.WaitGroup

	//for j := 0; j < len(lists); j++ {
	//	wg.Add(1)
	//	go worker(j, lists[j], &wg)
	//}

	//wg.Wait()

}
