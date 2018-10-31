package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		fileTable    []string
		mergeTable   []string
		layerNumber  int
		mergeSize    int
		outputNumber int
	)
	for i := 1; i < 12; i++ {
		fileName := "div/divFile" + strconv.Itoa(i) + ".txt"
		fileTable = append(fileTable, fileName)
	}
	start := time.Now()
	for len(fileTable) > 1 {
		outputNumber++
		layerNumber = int(math.Log2(float64(len(fileTable))))
		mergeSize = int(math.Pow(2, float64(layerNumber)))
		mergeTable, fileTable = fileTable[0:mergeSize], fileTable[mergeSize:]
		fileTable = append(fileTable, "div/output"+strconv.Itoa(outputNumber)+".txt")
		fmt.Println("fileTable:", fileTable)
		fmt.Println("mergeTable:", mergeTable)
		winnerTreeMerge(mergeTable, outputNumber)
	}
	end := time.Now()
	fmt.Println("time:", end.Sub(start).Seconds())
}

func winnerTreeMerge(fileTable []string, outputNumber int) {
	fmt.Println("Start winner merge !")

	var (
		numFile      = len(fileTable)
		mergeFile    = make([]*os.File, numFile)
		openErr      = make([]error, numFile)
		readErr      = make([]error, numFile)
		scanner      = make([]*bufio.Reader, numFile)
		countNilFile int
	)

	layer := make([][]string, int(math.Log2(float64(numFile))+1))
	for i := range layer {
		layer[i] = make([]string, int(math.Pow(2, float64(i))))
	}
	//來自哪個節點
	frombuf := make([][]int, int(math.Log2(float64(numFile))))
	for i := range frombuf {
		frombuf[i] = make([]int, int(math.Pow(2, float64(i))))
	}
	//比較node的大小
	compareNode := make([][]int, int(math.Log2(float64(numFile))+1))
	for i := range compareNode {
		compareNode[i] = make([]int, int(math.Pow(2, float64(i))))
	}

	for i := 0; i < numFile; i++ {
		mergeFile[i], openErr[i] = os.Open(fileTable[i])
		defer mergeFile[i].Close()
		scanner[i] = bufio.NewReader(mergeFile[i])
		layer[len(layer)-1][i], readErr[i] = scanner[i].ReadString('\n')
		compareNode[len(layer)-1][i] = strings.Count(layer[len(layer)-1][i], "")
	}
	outputName := "div/output" + strconv.Itoa(outputNumber) + ".txt"
	outputfile, outputerr := os.Create(outputName)
	if outputerr != nil {
		fmt.Println("Creat file error")
		os.Exit(1)
	}
	defer outputfile.Close()

	for countNilFile < numFile {
		for i := (len(layer) - 1); i > 0; i-- {
			for j := 0; j < int(math.Pow(2, float64(i))); j = j + 2 {
				if compareNode[i][j] <= compareNode[i][j+1] {
					if j != 0 {
						layer[i-1][j/2] = layer[i][j]
						compareNode[i-1][j/2] = compareNode[i][j]
						if i == len(layer)-1 {
							frombuf[i-1][j/2] = j
						} else {
							frombuf[i-1][j/2] = frombuf[i][j]
						}
					} else {
						layer[i-1][0] = layer[i][j]
						compareNode[i-1][0] = compareNode[i][j]
						if i == len(layer)-1 {
							frombuf[i-1][0] = j
						} else {
							frombuf[i-1][0] = frombuf[i][j]
						}
					}
				} else {
					if j != 0 {
						layer[i-1][j/2] = layer[i][j+1]
						compareNode[i-1][j/2] = compareNode[i][j+1]
						if i == len(layer)-1 {
							frombuf[i-1][j/2] = j + 1
						} else {
							frombuf[i-1][j/2] = frombuf[i][j+1]
						}
					} else {
						layer[i-1][0] = layer[i][j+1]
						compareNode[i-1][0] = compareNode[i][j+1]
						if i == len(layer)-1 {
							frombuf[i-1][0] = j + 1
						} else {
							frombuf[i-1][0] = frombuf[i][j+1]
						}
					}
				}
			}
		}
		outputfile.WriteString(fmt.Sprintf("%s", layer[0][0]))
		layer[len(layer)-1][frombuf[0][0]], readErr[frombuf[0][0]] = scanner[frombuf[0][0]].ReadString('\n')
		compareNode[len(layer)-1][frombuf[0][0]] = strings.Count(layer[len(layer)-1][frombuf[0][0]], "")
		if readErr[frombuf[0][0]] != nil {
			compareNode[len(layer)-1][frombuf[0][0]] = 10000
			countNilFile++
		}
	}

	//delete the file that have merged
	// for i := 0; i < numFile; i++ {
	// 	mergeFile[i].Close()
	// 	os.Remove(fileTable[i])
	// }
}
