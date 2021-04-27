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

func makeIntitialLists(maxNumberToGen int, listSizeCap int, moduloMax int, nQuasiprime int, outputDir string) map[int]QuasiprimeList {
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
		quasiprimeList.outFileName = fmt.Sprintf("%s/%v-quasiprimes.moduloMax%v.part.%016d.txt", outputDir, nQuasiprime, moduloMax, i)
		quasiprimeList.moduloMax = moduloMax
		quasiprimeList.nQuasiprime = nQuasiprime

		moduloDataList := make(map[int]map[int]moduloData, moduloMax-1)
		for m := 2; m <= moduloMax; m++ {
			moduloDataList[m] = make(map[int]moduloData, m)
			for p := 0; p < m; p++ {
				moduloDataList[m][p] = moduloData{}
			}
		}
		quasiprimeList.moduloDataList = moduloDataList

		pairedModuloDataList := make(map[int]map[int]map[int]moduloData, moduloMax-1)
		for m := 2; m <= moduloMax; m++ {
			pairedModuloDataList[m] = make(map[int]map[int]moduloData, m)
			for a := 0; a < m; a++ {
				pairedModuloDataList[m][a] = make(map[int]moduloData, m)
				for b := 0; b < m; b++ {
					pairedModuloDataList[m][a][b] = moduloData{}
				}
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
