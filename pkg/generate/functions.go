package generate

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// vPrint message is the verbosity is greater than or equal to the minimum
func vPrint(verbosity int, minimum int, message string, args ...interface{}) {
	if verbosity >= minimum {
		fmt.Printf(message, args...)
	}
}

func makeIntitialLists(maxNumberToGen int, listSizeCap int, modulo int, outputDir string) map[int]QuasiprimeList {
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

		pairedModuloDataList := make(map[int]map[int]moduloData, modulo)
		for a := 0; a < modulo; a++ {
			pairedModuloDataList[a] = make(map[int]moduloData, modulo)
			for b := 0; b < modulo; b++ {
				pairedModuloDataList[a][b] = moduloData{}
			}
		}
		quasiprimeList.pairedModuloDataList = pairedModuloDataList

		lists[i] = quasiprimeList
	}

	return lists
}

func makeMasterPrimeList(maxNumberToGen int, primeSourceFile string) map[int]int {
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

	return masterPrimeList
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
