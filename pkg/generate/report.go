package generate

import "fmt"

func report(maxNumberToGen int, listSizeCap int, moduloMax int, nQuasiprime int, outputDir string, quasiprimeListChannel chan QuasiprimeList) {
	var completeQuasiprimeList QuasiprimeList

	completeQuasiprimeList.moduloMax = moduloMax
	completeQuasiprimeList.nQuasiprime = nQuasiprime
	completeQuasiprimeList.outFileName = fmt.Sprintf("%s/%v-quasiprimes.moduloMax%v.numberMax%v.complete_report.txt", outputDir, nQuasiprime, moduloMax, maxNumberToGen)
	completeQuasiprimeList.minIntegerChecked = 0
	completeQuasiprimeList.maxIntegerChecked = maxNumberToGen
	completeQuasiprimeList.numIntergersChecked = maxNumberToGen + 1

	coll := make(map[int]map[int]int, len(quasiprimeListChannel))
	mins := make(map[int]int, len(quasiprimeListChannel))
	order := make(map[int]int, len(quasiprimeListChannel))
	i := 0
	for quasiprimeList := range quasiprimeListChannel {
		completeQuasiprimeList.numQuasiprimes += quasiprimeList.numQuasiprimes
		completeQuasiprimeList.quasiprimeGenerationTime += quasiprimeList.quasiprimeGenerationTime

		if len(completeQuasiprimeList.primes) < len(quasiprimeList.primes) {
			completeQuasiprimeList.primes = quasiprimeList.primes
		}

		coll[i] = quasiprimeList.quasiprimes
		mins[i] = quasiprimeList.quasiprimes[0]
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

	completeQuasiprimes := make(map[int]int, completeQuasiprimeList.numQuasiprimes)
	k := 0
	for i := 0; i < len(order); i++ {
		for j := 0; j < len(coll[order[i]]); j++ {
			completeQuasiprimes[k] = coll[order[i]][j]
			k++
		}
	}

	completeQuasiprimeList.quasiprimes = completeQuasiprimes
	completeQuasiprimeList.minQuasiprime = completeQuasiprimeList.quasiprimes[0]
	completeQuasiprimeList.maxQuasiprime = completeQuasiprimeList.quasiprimes[completeQuasiprimeList.numQuasiprimes-1]

	moduloDataList := make(map[int]map[int]moduloData, moduloMax-1)
	for m := 2; m <= moduloMax; m++ {
		moduloDataList[m] = make(map[int]moduloData, m)
		for p := 0; p < m; p++ {
			quantity := 0
			for _, quasiprime := range completeQuasiprimeList.quasiprimes {
				if quasiprime%m == p {
					quantity++
				}
			}
			percentage := float64(quantity) / float64(completeQuasiprimeList.numQuasiprimes)

			moduloDataList[m][p] = moduloData{quantity, percentage}
		}
	}
	completeQuasiprimeList.moduloDataList = moduloDataList

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
	completeQuasiprimeList.pairedModuloDataList = pairedModuloDataList

	for m := range completeQuasiprimeList.pairedModuloDataList {
		totalPairs := 0
		for i := 0; i < len(completeQuasiprimeList.quasiprimes)-1; i++ {
			q, n := completeQuasiprimeList.quasiprimes[i], completeQuasiprimeList.quasiprimes[i+1]

			completeQuasiprimeList.pairedModuloDataList[m][q%m][n%m] =
				moduloData{completeQuasiprimeList.pairedModuloDataList[m][q%m][n%m].quantity + 1, 0.0}

			totalPairs++
		}

		for a := 0; a < len(completeQuasiprimeList.pairedModuloDataList[m]); a++ {
			for b := 0; b < len(completeQuasiprimeList.pairedModuloDataList[m][a]); b++ {
				completeQuasiprimeList.pairedModuloDataList[m][a][b] = moduloData{completeQuasiprimeList.pairedModuloDataList[m][a][b].quantity,
					float64(completeQuasiprimeList.pairedModuloDataList[m][a][b].quantity) / float64(totalPairs)}
			}
		}
	}

	completeQuasiprimeList.writeToFile(false, true)
}
